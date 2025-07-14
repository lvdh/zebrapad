package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockPrinter struct {
	Printed []byte
	Err     error
}

func (m *mockPrinter) Print(data []byte) error {
	m.Printed = data
	return m.Err
}

func TestIndexRoute(t *testing.T) {
	srv := NewHTTPServer(&mockPrinter{})
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		t.Error("expected non-empty body for /")
	}
}

func TestAssetsRoute_NotFound(t *testing.T) {
	srv := NewHTTPServer(&mockPrinter{})
	req := httptest.NewRequest("GET", "/assets/doesnotexist.js", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404 Not Found, got %d", resp.StatusCode)
	}
}

func TestPrintRoute_Success(t *testing.T) {
	mp := &mockPrinter{}
	srv := NewHTTPServer(mp)
	payload := map[string]string{"zpl": "^XA^FO50,50^FDTest^FS^XZ"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/print", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
	if string(mp.Printed) != payload["zpl"] {
		t.Errorf("expected printer to receive %q, got %q", payload["zpl"], string(mp.Printed))
	}
}

func TestPrintRoute_BadRequest(t *testing.T) {
	srv := NewHTTPServer(&mockPrinter{})
	req := httptest.NewRequest("POST", "/print", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", resp.StatusCode)
	}
}

func TestPrintRoute_MethodNotAllowed(t *testing.T) {
	srv := NewHTTPServer(&mockPrinter{})
	req := httptest.NewRequest("GET", "/print", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method Not Allowed, got %d", resp.StatusCode)
	}
}
