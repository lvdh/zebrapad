
//go:build !windows

package printer

// DiscoverLocalPrinters is a stub for non-Windows platforms
func DiscoverLocalPrinters() ([]DiscoveredPrinter, error) {
	// Return empty slice on non-Windows platforms
	return []DiscoveredPrinter{}, nil
}
