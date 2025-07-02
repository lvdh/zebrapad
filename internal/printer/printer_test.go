package printer

import (
	"net"
	"testing"
	"time"
)

func TestSendZPLToPrinter(t *testing.T) {
	// Start a mock TCP server
	ln, err := net.Listen("tcp", "127.0.0.1:19100")
	if err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer ln.Close()

	PrinterAddress = "127.0.0.1:19100"

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("mock server accept error: %v", err)
			return
		}
		defer conn.Close()
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Errorf("mock server read error: %v", err)
			return
		}
		got := string(buf[:n])
		want := "^XA^FO50,50^ADN,36,20^FDTest Label^FS^XZ"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	}()

	// Give the goroutine a moment to start
	time.Sleep(100 * time.Millisecond)

	err = SendZPLToPrinter("^XA^FO50,50^ADN,36,20^FDTest Label^FS^XZ")
	if err != nil {
		t.Fatalf("SendZPLToPrinter failed: %v", err)
	}
}
