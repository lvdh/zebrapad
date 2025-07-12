package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lvdh/zebrapad/internal/printer"
)

type mockPrinter struct {
	received []byte
	err      error
}

func (m *mockPrinter) Print(data []byte) error {
	m.received = data
	return m.err
}

func setupRouter(p printer.Printer) *gin.Engine {
	gin.SetMode(gin.TestMode)
	return NewHTTPServer(p)
}

func TestPrintEndpoint_Success(t *testing.T) {
	mp := &mockPrinter{}
	router := setupRouter(mp)

	body := `{"zpl":"^XA^FO50,50^ADN,36,20^FDHello^FS^XZ"}`
	req, _ := http.NewRequest("POST", "/print", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if string(mp.received) != "^XA^FO50,50^ADN,36,20^FDHello^FS^XZ" {
		t.Errorf("printer received wrong data: %q", mp.received)
	}
}

func TestPrintEndpoint_BadRequest(t *testing.T) {
	mp := &mockPrinter{}
	router := setupRouter(mp)

	body := `{}`
	req, _ := http.NewRequest("POST", "/print", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestPrintEndpoint_PrinterError(t *testing.T) {
	mp := &mockPrinter{err: errors.New("printer unavailable")}
	router := setupRouter(mp)

	body := `{"zpl":"test"}`
	req, _ := http.NewRequest("POST", "/print", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
