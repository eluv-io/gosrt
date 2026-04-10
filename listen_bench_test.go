package srt

import (
	"net"
	"testing"

	"github.com/datarhei/gosrt/packet"
)

// BenchmarkListenerRecvDump benchmarks the packet:recv:dump log call in the listener reader loop —
// exercised once per incoming UDP packet. With logging disabled the logEnabled guard prevents the
// message closure from being heap-allocated, dropping allocs/op to zero.
func BenchmarkListenerRecvDump(b *testing.B) {
	addr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		ln := &listener{config: Config{Logger: logger}}
		p := packet.NewPacket(addr)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if ln.logEnabled("packet:recv:dump") {
				ln.log("packet:recv:dump", func() string { return p.Dump() })
			}
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"packet:recv:dump"})) })
}

// BenchmarkListenerSend benchmarks listener.send() — the function called to write every outgoing
// packet to the wire. Uses a real UDP socket so the full code path including Marshal is exercised.
// With logging disabled the logEnabled guard on packet:send:dump eliminates the closure allocation.
func BenchmarkListenerSend(b *testing.B) {
	pc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		b.Fatal(err)
	}
	defer pc.Close()

	dst := pc.LocalAddr().(*net.UDPAddr) // send to self; datagrams are silently dropped when recv buf fills

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		ln := &listener{
			pc:     pc,
			config: Config{Logger: logger},
		}
		p := packet.NewPacket(dst) // data packet — not decommissioned by send, safe to reuse
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ln.send(p)
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"packet:send:dump"})) })
}
