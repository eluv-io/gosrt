package srt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// acceptOne listens on the test address and returns the listener with a channel that receives the server side of the
// first accepted connection.
func acceptOne(t *testing.T) (Listener, chan Conn) {
	ln, err := Listen("srt", "127.0.0.1:6003", DefaultConfig())
	require.NoError(t, err)

	accepted := make(chan Conn, 1)
	go func() {
		req, err := ln.Accept2()
		if err != nil {
			return
		}
		conn, err := req.Accept()
		if err != nil {
			return
		}
		accepted <- conn
	}()

	return ln, accepted
}

func requireDone(t *testing.T, conn Conn, msg string) {
	select {
	case <-conn.Done():
	case <-time.After(3 * time.Second):
		require.Fail(t, msg)
	}
}

// TestConnDonePeerClose checks that the Done channel of a connection closes when the peer closes the connection,
// without any data passing.
func TestConnDonePeerClose(t *testing.T) {
	ln, accepted := acceptOne(t)
	defer ln.Close()

	client, err := Dial("srt", "127.0.0.1:6003", DefaultConfig())
	require.NoError(t, err)
	server := <-accepted

	select {
	case <-server.Done():
		require.Fail(t, "the connection is open")
	default:
	}

	require.NoError(t, client.Close())
	requireDone(t, client, "the client's own close closes its done channel")
	requireDone(t, server, "the peer's close closes the done channel")
}

// TestConnDoneListenerClose checks that closing a listener closes the Done channels of its connections.
func TestConnDoneListenerClose(t *testing.T) {
	ln, accepted := acceptOne(t)

	client, err := Dial("srt", "127.0.0.1:6003", DefaultConfig())
	require.NoError(t, err)
	defer func() { _ = client.Close() }()
	server := <-accepted

	ln.Close()
	requireDone(t, server, "closing the listener closes the done channel")
}
