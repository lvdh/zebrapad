package printer

// DiscoveredPrinter represents a printer found during discovery
type DiscoveredPrinter struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
	Type string `json:"type"` // "local", "network", or "manual"
}
