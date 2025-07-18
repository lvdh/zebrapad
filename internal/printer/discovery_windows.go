//go:build windows

package printer

import (
	"fmt"
	"log"

	"github.com/StackExchange/wmi"
)

// Win32_Printer represents a Windows printer from WMI
type Win32_Printer struct {
	Name       string
	ShareName  string
	PortName   string
	DriverName string
	Comment    string
	Location   string
	Local      bool
}

// DiscoverLocalPrinters discovers local Windows printers using WMI
func DiscoverLocalPrinters() ([]DiscoveredPrinter, error) {
	var printers []Win32_Printer
	var discovered []DiscoveredPrinter

	// Query WMI for all printers
	q := wmi.CreateQuery(&printers, "")
	err := wmi.Query(q, &printers)
	if err != nil {
		log.Printf("WMI query error: %v", err)
		return discovered, err
	}

	for _, p := range printers {
		// Skip fax and XPS printers, focus on physical printers
		if isPhysicalPrinter(p.DriverName, p.PortName) {
			name := p.Name
			if p.Location != "" {
				name = fmt.Sprintf("%s (%s)", p.Name, p.Location)
			}

			discovered = append(discovered, DiscoveredPrinter{
				Name: name,
				URI:  fmt.Sprintf("winspool://%s", p.Name),
				Type: "local",
			})
		}
	}

	return discovered, nil
}

// isPhysicalPrinter filters out virtual printers like fax, XPS, etc.
func isPhysicalPrinter(driverName, portName string) bool {
	// Common virtual printer patterns to exclude
	virtualPatterns := []string{
		"Microsoft XPS",
		"Microsoft Print to PDF",
		"Fax",
		"OneNote",
		"Send To OneNote",
	}

	for _, pattern := range virtualPatterns {
		if contains(driverName, pattern) || contains(portName, pattern) {
			return false
		}
	}

	// Include USB, parallel, and network ports
	physicalPorts := []string{"USB", "LPT", "COM", "IP_", "WSD"}
	for _, port := range physicalPorts {
		if contains(portName, port) {
			return true
		}
	}

	return true
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					findInString(s, substr)))
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
