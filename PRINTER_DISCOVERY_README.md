# Zebrapad Printer Auto-Discovery Extension

This document describes the printer auto-discovery functionality that has been added to the zebrapad project.

## Overview

The enhanced zebrapad now includes automatic printer discovery capabilities that work on Windows systems, with both local printer enumeration and network scanning for Zebra printers.

## Features Added

### Backend Extensions (Go)

1. **Windows Local Printer Discovery** (`internal/printer/discovery_windows.go`)
   - Uses Windows WMI (Windows Management Instrumentation) to enumerate local printers
   - Filters out virtual printers (PDF, XPS, Fax, etc.)
   - Returns printer names with location information when available
   - Only compiles on Windows systems (build tag: `//go:build windows`)

2. **Network Printer Scanning** (`internal/printer/network_scan.go`)
   - Scans local subnet for Zebra printers on common ports (9100, 6101, 9101, 9102, 9103)
   - Uses concurrent goroutines for efficient scanning (50 workers by default)
   - Attempts to identify printers using ZPL `~HI` command
   - Provides fallback naming for unidentified network printers
   - Works on all platforms

3. **API Endpoint** (`internal/server/api.go`)
   - New `/api/printers` GET endpoint returns discovered printers in JSON format
   - Combines local and network discovery results
   - Always includes a manual TCP entry (`:9010`)
   - Graceful error handling - continues operation even if discovery fails

4. **Cross-Platform Support** (`internal/printer/discovery_stub.go`)
   - Stub implementation for non-Windows platforms
   - Ensures the application compiles and runs on Linux/macOS

### Frontend Extensions (HTML/JS)

1. **Printer Selection UI** (updated `internal/server/web/index.html`)
   - Added printer dropdown at the top of the page
   - Manual refresh button with loading states
   - Clean Bulma CSS styling that matches existing design

2. **JavaScript Integration** (`internal/server/web/js/printers.js`)
   - Automatic printer discovery on page load
   - Manual refresh functionality
   - Updates dropdown with discovered printers
   - Visual indicators for printer types (🖨️ for local, 🌐 for network)
   - Error handling and user feedback

3. **Form Integration** (updated `internal/server/web/js/forms.js`)
   - Modified print functions to include selected printer information
   - Maintains backward compatibility with existing functionality

## Technical Implementation

### Dependencies Added
- `github.com/StackExchange/wmi` - For Windows WMI queries
- `github.com/go-ole/go-ole` - COM/OLE support (dependency of WMI)
- `golang.org/x/sys` - System-specific functionality

### Network Scanning Details
- **Ports Scanned**: 9100 (primary), 6101, 9101, 9102, 9103
- **Timeout**: 2 seconds connection timeout, 3 seconds read/write timeout
- **Concurrency**: 50 concurrent workers for efficient scanning
- **Subnet Detection**: Automatically detects local subnet from network interfaces
- **ZPL Identification**: Uses `~HI` command to query printer model information

### Windows Printer Discovery Details
- **WMI Class**: `Win32_Printer`
- **Filtering**: Excludes virtual printers (PDF, XPS, Fax, OneNote, etc.)
- **Port Types**: Includes USB, LPT, COM, IP_, WSD ports
- **Naming**: Combines printer name with location when available

## API Response Format

```json
{
  "printers": [
    {
      "name": "Manual TCP (:9010)",
      "uri": ":9010",
      "type": "manual"
    },
    {
      "name": "Zebra ZT410 (192.168.1.100)",
      "uri": "192.168.1.100:9100",
      "type": "network"
    },
    {
      "name": "Zebra GK420t (Office)",
      "uri": "winspool://Zebra GK420t",
      "type": "local"
    }
  ],
  "error": "Network discovery failed: timeout"
}
```

## Usage Instructions

### Running on Windows
1. Use the provided `zebrapad.exe` executable
2. Local printer discovery will automatically work
3. Network scanning will discover Zebra printers on the local subnet

### Running on Linux/macOS
1. Use the `zebrapad` executable
2. Only network scanning will be available (local discovery is Windows-only)
3. Manual TCP option is always available

### Web Interface
1. Open browser to `http://localhost:8080`
2. Select printer from dropdown at the top of the page
3. Click "Refresh" to re-scan for printers
4. Use existing Library Card or ZPL Input functionality as before

## Build Instructions

### For Current Platform
```bash
go build ./cmd/zebrapad
```

### For Windows (Cross-compile)
```bash
GOOS=windows GOARCH=amd64 go build -o zebrapad.exe ./cmd/zebrapad
```

### For Linux (Cross-compile)
```bash
GOOS=linux GOARCH=amd64 go build -o zebrapad-linux ./cmd/zebrapad
```

## Security Considerations

- Network scanning is limited to local subnet only
- WMI queries require appropriate Windows permissions
- No sensitive information is exposed through the API
- Timeouts prevent hanging on unresponsive network addresses

## Troubleshooting

### Windows Local Discovery Issues
- Ensure the application has appropriate permissions
- Check that WMI service is running
- Verify printers are properly installed in Windows

### Network Discovery Issues
- Check firewall settings (ports 9100, 6101, etc.)
- Ensure printers are on the same subnet
- Verify printer network configuration
- Check that printers support ZPL protocol

### General Issues
- Check console/logs for error messages
- Verify network connectivity
- Ensure Go dependencies are properly installed

## Future Enhancements

Potential improvements for future versions:
1. Support for additional printer protocols (IPP, SNMP)
2. Printer status monitoring
3. Custom port range configuration
4. Saved printer configurations
5. Printer capability detection
6. Integration with printer-specific settings

## Files Modified/Added

### New Files
- `internal/printer/types.go` - Common type definitions
- `internal/printer/discovery_windows.go` - Windows printer discovery
- `internal/printer/network_scan.go` - Network scanning functionality
- `internal/printer/discovery_stub.go` - Non-Windows stub
- `internal/server/api.go` - API endpoint handler
- `internal/server/web/js/printers.js` - Frontend printer management

### Modified Files
- `internal/server/server.go` - Added API endpoint registration
- `internal/server/web/index.html` - Added printer selection UI
- `internal/server/web/js/forms.js` - Updated to use selected printer
- `go.mod` - Added required dependencies

This implementation provides a solid foundation for printer auto-discovery while maintaining the existing zebrapad functionality and ensuring cross-platform compatibility.
