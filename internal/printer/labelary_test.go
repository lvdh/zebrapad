package printer

//import (
//	"fmt"
//	"net"
//	"testing"
//	"time"
//)
//
//func startFakeLabelaryPrinter(t *testing.T, received chan<- []byte) (addr string, stop func()) {
//	ln, err := net.Listen("tcp", "127.0.0.1:0")
//	if err != nil {
//		t.Fatalf("failed to start fake printer: %v", err)
//	}
//	go func() {
//		conn, err := ln.Accept()
//		if err != nil {
//			return
//		}
//		defer conn.Close()
//		buf := make([]byte, 1024)
//		n, _ := conn.Read(buf)
//		received <- buf[:n]
//	}()
//	return ln.Addr().String(), func() { ln.Close() }
//}
//
//func TestLabelaryPrinter_Print_Success(t *testing.T) {
//	received := make(chan []byte, 1)
//	addr, stop := startFakeLabelaryPrinter(t, received)
//	defer stop()
//
//	p := &LabelaryPrinter{
//		// "http://api.labelary.com/v1/printers"
//		url:     fmt.Sprintf("%s/v1/printers", addr),
//		density: "8dpmm",
//		width:   2.25,
//		height:  0.75,
//	}
//
//	data := []byte("test label")
//	err := p.Print(data)
//	if err != nil {
//		t.Fatalf("Print failed: %v", err)
//	}
//
//	select {
//	case got := <-received:
//		if string(got) != string(data) {
//			t.Errorf("expected %q, got %q", data, got)
//		}
//	case <-time.After(time.Second):
//		t.Fatal("timeout waiting for printer data")
//	}
//}
//
//func TestLabelaryPrinter_Print_ConnectionError(t *testing.T) {
//	received := make(chan []byte, 1)
//	addr, stop := startFakeLabelaryPrinter(t, received)
//	defer stop()
//
//	p := &LabelaryPrinter{
//		// "http://api.labelary.com/v1/printers"
//		url:     fmt.Sprintf("%s/v1/printers", addr),
//		density: "8dpmm",
//		width:   2.25,
//		height:  0.75,
//	}
//	err := p.Print([]byte("fail"))
//	if err == nil {
//		t.Fatal("expected error, got nil")
//	}
//}
