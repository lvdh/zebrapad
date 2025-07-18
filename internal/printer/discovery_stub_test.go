package printer

import (
	"testing"
)

func TestDiscoverLocalPrinters_Stub(t *testing.T) {
	printers, err := DiscoverLocalPrinters()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(printers) != 0 {
		t.Errorf("expected empty slice, got %v", printers)
	}
}
