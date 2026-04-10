package srt

import (
	"net"
	"testing"

	"github.com/datarhei/gosrt/packet"
)

// BenchmarkDialerRecvDump benchmarks the packet:recv:dump log call in the dialer reader loop —
// exercised once per incoming UDP packet. With logging disabled the logEnabled guard prevents the
// message closure from being heap-allocated, dropping allocs/op to zero.
func BenchmarkDialerRecvDump(b *testing.B) {
	addr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		dl := &dialer{config: Config{Logger: logger}}
		p := packet.NewPacket(addr)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if dl.logEnabled("packet:recv:dump") {
				dl.log("packet:recv:dump", func() string { return p.Dump() })
			}
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"packet:recv:dump"})) })
}

// BenchmarkDialerSend benchmarks dialer.send() — the function called to write every outgoing
// packet to the wire. Uses a real connected UDP socket so the full code path including Marshal is
// exercised. With logging disabled the logEnabled guard on packet:send:dump eliminates the closure
// allocation.
func BenchmarkDialerSend(b *testing.B) {
	// The dialer uses a connected UDP socket (pc.Write, not WriteTo).
	srv, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		b.Fatal(err)
	}
	defer srv.Close()

	cli, err := net.DialUDP("udp", nil, srv.LocalAddr().(*net.UDPAddr))
	if err != nil {
		b.Fatal(err)
	}
	defer cli.Close()

	run := func(b *testing.B, logger Logger) {
		b.Helper()
		dl := &dialer{
			pc:     cli,
			config: Config{Logger: logger},
		}
		p := packet.NewPacket(cli.RemoteAddr()) // data packet — not decommissioned by send, safe to reuse
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dl.send(p)
		}
	}

	b.Run("logging_disabled", func(b *testing.B) { run(b, NewLogger(nil)) })
	b.Run("logging_enabled", func(b *testing.B) { run(b, NewLogger([]string{"packet:send:dump"})) })
}
