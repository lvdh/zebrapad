// internal/printer/discovery_windows_test.go
//go:build windows

package printer

import (
	"testing"
)

func TestIsPhysicalPrinter_True(t *testing.T) {
	// Should return true for physical port
	if !isPhysicalPrinter("Zebra Driver", "USB001") {
		t.Error("expected true for USB port")
	}
}

func TestIsPhysicalPrinter_FalseVirtual(t *testing.T) {
	// Should return false for virtual printer
	if isPhysicalPrinter("Microsoft XPS", "PORTPROMPT:") {
		t.Error("expected false for virtual printer")
	}
	if isPhysicalPrinter("Zebra Driver", "Microsoft Print to PDF") {
		t.Error("expected false for virtual port")
	}
}

func TestContains(t *testing.T) {
	if !contains("Microsoft XPS Document Writer", "XPS") {
		t.Error("expected true for substring")
	}
	if contains("Zebra", "Canon") {
		t.Error("expected false for non-matching substring")
	}
}

func TestFindInString(t *testing.T) {
	if !findInString("abcdef", "cde") {
		t.Error("expected true for substring")
	}
	if findInString("abcdef", "xyz") {
		t.Error("expected false for missing substring")
	}
}

// Integration test for DiscoverLocalPrinters (mocking WMI is non-trivial, so just check error path)
func TestDiscoverLocalPrinters_Error(t *testing.T) {
	// Simulate WMI error by passing an invalid query (not possible directly, so just check that it returns slice/error)
	printers, err := DiscoverLocalPrinters()
	// Accept either: error (if WMI fails), or printers (if running on Windows with printers)
	if err != nil && len(printers) != 0 {
		t.Errorf("expected empty slice on error, got %v", printers)
	}
}
