package printer

import (
	"net"
	"testing"
)

func startFakePrinter(t *testing.T, received chan<- []byte) (addr string, stop func()) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start fake printer: %v", err)
	}
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		received <- buf[:n]
	}()
	return ln.Addr().String(), func() { ln.Close() }
}
