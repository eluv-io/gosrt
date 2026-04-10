package srt

import (
	"net"
	"testing"
	"time"

	"github.com/datarhei/gosrt/circular"
	"github.com/datarhei/gosrt/congestion/live"
	"github.com/datarhei/gosrt/packet"
)

// BenchmarkPopDataPacket benchmarks pop() on the data-packet path — the hot path exercised once
// per sent packet by the congestion sender. With logging disabled (the production default), the
// logEnabled guard prevents the message closure from being heap-allocated, dropping allocs/op to
// zero.
func BenchmarkPopDataPacket(b *testing.B) {
	addr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		conn := &srtConn{
			localAddr:  addr,
			remoteAddr: addr,
			logger:     logger,
			onSend:     func(p packet.Packet) {}, // data packets are not decommissioned by onSend
		}
		p := packet.NewPacket(addr) // pre-allocated; reused across iterations
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			conn.pop(p)
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"data:send:dump"})) })
}

// BenchmarkSendACK benchmarks sendACK() — called roughly once per ACK interval (~100/s per
// connection). lite=true is used to avoid map-growth allocations in ackNumbers that would obscure
// the closure allocation signal; the guarded log calls exist in both lite and full-ACK paths.
// With logging disabled the 2 allocs/op that remain are pre-existing and unrelated to logging: CIFACK
// escapes to the heap via the CIF interface in MarshalCIF, and its internal [28]byte marshal buffer
// escapes through io.Writer. With logging enabled those 2 grow to ~28 (closures + p.Dump output).
func BenchmarkSendACK(b *testing.B) {
	addr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}
	seq := circular.New(0, packet.MAX_SEQUENCENUMBER)

	newConn := func(logger Logger) *srtConn {
		return &srtConn{
			localAddr:     addr,
			remoteAddr:    addr,
			logger:        logger,
			onSend:        func(p packet.Packet) { p.Decommission() },
			nextACKNumber: circular.New(1, packet.MAX_TIMESTAMP),
			ackNumbers:    make(map[uint32]time.Time),
			config:        DefaultConfig(),
			recv: live.NewReceiver(live.ReceiveConfig{
				InitialSequenceNumber: seq,
				PeriodicACKInterval:   10_000,
				PeriodicNAKInterval:   20_000,
				OnSendACK:             func(circular.Number, bool) {},
				OnSendNAK:             func([]circular.Number) {},
				OnDeliver:             func(packet.Packet) {},
			}),
			rtt: rtt{rtt: 100_000, rttVar: 50_000},
		}
	}

	b.Run("logging_disabled", func(b *testing.B) {
		conn := newConn(NewLogger(nil))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			conn.sendACK(seq, true) // lite=true: skips ackNumbers map write
		}
	})
	b.Run("logging_enabled", func(b *testing.B) {
		conn := newConn(NewLogger([]string{"control:send:ACK"}))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			conn.sendACK(seq, true)
		}
	})
}

// BenchmarkSendACKACK benchmarks sendACKACK() — called once per full ACK received (~100/s).
func BenchmarkSendACKACK(b *testing.B) {
	addr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		conn := &srtConn{
			localAddr:  addr,
			remoteAddr: addr,
			logger:     logger,
			onSend:     func(p packet.Packet) { p.Decommission() },
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			conn.sendACKACK(42)
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"control:send:ACKACK:dump"})) })
}
