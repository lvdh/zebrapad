package printer

import (
	"net"
	"testing"
	"time"
)

func startFakeTCPPrinter(t *testing.T, received chan<- []byte) (addr string, stop func()) {
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

func TestTCPPrinter_Print_Success(t *testing.T) {
	received := make(chan []byte, 1)
	addr, stop := startFakeTCPPrinter(t, received)
	defer stop()

	p := &TCPPrinter{Network: "tcp", Address: addr}
	data := []byte("test label")
	err := p.Print(data)
	if err != nil {
		t.Fatalf("Print failed: %v", err)
	}

	select {
	case got := <-received:
		if string(got) != string(data) {
			t.Errorf("expected %q, got %q", data, got)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for printer data")
	}
}

func TestTCPPrinter_Print_ConnectionError(t *testing.T) {
	p := &TCPPrinter{Network: "tcp", Address: "127.0.0.1:65000"} // unused port
	err := p.Print([]byte("fail"))
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
