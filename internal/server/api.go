
package server

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/lvdh/zebrapad/internal/printer"
)

// PrintersResponse represents the API response for printer discovery
type PrintersResponse struct {
	Printers []printer.DiscoveredPrinter `json:"printers"`
	Error    string                      `json:"error,omitempty"`
}

// handlePrintersAPI handles the /api/printers endpoint
func handlePrintersAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var allPrinters []printer.DiscoveredPrinter
	var errors []string

	// Always add the dummy TCP printer entry
	allPrinters = append(allPrinters, printer.DiscoveredPrinter{
		Name: "Manual TCP (:9010)",
		URI:  ":9010",
		Type: "manual",
	})

	// Discover network printers (works on all platforms)
	networkPrinters, err := printer.DiscoverNetworkPrinters()
	if err != nil {
		errors = append(errors, "Network discovery failed: "+err.Error())
	} else {
		allPrinters = append(allPrinters, networkPrinters...)
	}

	// Discover local printers (Windows only)
	if runtime.GOOS == "windows" {
		localPrinters, err := printer.DiscoverLocalPrinters()
		if err != nil {
			errors = append(errors, "Local discovery failed: "+err.Error())
		} else {
			allPrinters = append(allPrinters, localPrinters...)
		}
	}

	response := PrintersResponse{
		Printers: allPrinters,
	}

	// Include errors in response if any occurred
	if len(errors) > 0 {
		errorMsg := ""
		for i, err := range errors {
			if i > 0 {
				errorMsg += "; "
			}
			errorMsg += err
		}
		response.Error = errorMsg
	}

	json.NewEncoder(w).Encode(response)
}
