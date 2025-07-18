
package printer

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Common Zebra printer ports
var zebraPorts = []int{9100, 6101, 9101, 9102, 9103}

// DiscoverNetworkPrinters scans the local subnet for Zebra printers
func DiscoverNetworkPrinters() ([]DiscoveredPrinter, error) {
	var discovered []DiscoveredPrinter
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Get local subnet
	subnet, err := getLocalSubnet()
	if err != nil {
		return discovered, fmt.Errorf("failed to get local subnet: %w", err)
	}

	// Generate IP addresses to scan
	ips := generateSubnetIPs(subnet)

	// Channel for work distribution
	jobs := make(chan string, len(ips)*len(zebraPorts))
	results := make(chan DiscoveredPrinter, len(ips)*len(zebraPorts))

	// Start worker goroutines
	numWorkers := 50
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				parts := strings.Split(job, ":")
				if len(parts) != 2 {
					continue
				}
				ip := parts[0]
				port := parts[1]

				if printer := scanPrinter(ip, port); printer != nil {
					results <- *printer
				}
			}
		}()
	}

	// Send jobs
	go func() {
		for _, ip := range ips {
			for _, port := range zebraPorts {
				jobs <- fmt.Sprintf("%s:%d", ip, port)
			}
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Gather discovered printers
	for printer := range results {
		mu.Lock()
		discovered = append(discovered, printer)
		mu.Unlock()
	}

	return discovered, nil
}

// scanPrinter attempts to connect to a printer and identify it
func scanPrinter(ip, port string) *DiscoveredPrinter {
	address := fmt.Sprintf("%s:%s", ip, port)
	
	// Try to connect with timeout
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return nil
	}
	defer conn.Close()

	// Set read/write timeouts
	conn.SetDeadline(time.Now().Add(3 * time.Second))

	// Try to identify the printer using ZPL ~HI command
	name := identifyPrinter(conn, ip)
	
	return &DiscoveredPrinter{
		Name: name,
		URI:  address,
		Type: "network",
	}
}

// identifyPrinter sends ZPL identification command and parses response
func identifyPrinter(conn net.Conn, ip string) string {
	// Send ~HI command to get printer info
	_, err := conn.Write([]byte("~HI\n"))
	if err != nil {
		return fmt.Sprintf("Network Printer at %s", ip)
	}

	// Read response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil || n == 0 {
		return fmt.Sprintf("Network Printer at %s", ip)
	}

	response := string(buffer[:n])
	
	// Parse the response to extract printer model
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "Firmware") && 
		   !strings.HasPrefix(line, "Memory") && !strings.HasPrefix(line, "Printer Language") {
			// This is likely the printer model name
			if strings.Contains(line, "Zebra") || strings.Contains(line, "ZT") || 
			   strings.Contains(line, "ZD") || strings.Contains(line, "GK") ||
			   strings.Contains(line, "GX") {
				return fmt.Sprintf("%s (%s)", line, ip)
			}
		}
	}

	return fmt.Sprintf("Zebra Printer at %s", ip)
}

// getLocalSubnet gets the local subnet for scanning
func getLocalSubnet() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// Return the network address (e.g., "192.168.1.0/24")
				return ipnet.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no suitable network interface found")
}

// generateSubnetIPs generates all IP addresses in a subnet
func generateSubnetIPs(subnet string) []string {
	var ips []string
	
	_, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return ips
	}

	// Generate IPs in the subnet (skip network and broadcast addresses)
	ip := ipnet.IP.Mask(ipnet.Mask)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		// Skip network address (.0) and broadcast address (.255)
		if ip[3] != 0 && ip[3] != 255 {
			ips = append(ips, ip.String())
		}
		// Limit to reasonable range to avoid scanning entire /16 networks
		if len(ips) >= 254 {
			break
		}
	}

	return ips
}

// inc increments an IP address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
