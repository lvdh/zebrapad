// static/js/printers.js

// Printer discovery and management functionality

let discoveredPrinters = [];

// Fetch discovered printers from the API
async function fetchPrinters() {
    try {
        const response = await fetch('/api/printers');
        const data = await response.json();
        
        if (response.ok) {
            discoveredPrinters = data.printers || [];
            updatePrinterDropdown();
            
            // Show any errors in console but don't block UI
            if (data.error) {
                console.warn('Printer discovery warnings:', data.error);
            }
        } else {
            console.error('Failed to fetch printers:', data.error);
            showPrinterError('Failed to discover printers');
        }
    } catch (error) {
        console.error('Error fetching printers:', error);
        showPrinterError('Network error while discovering printers');
    }
}

// Update the printer dropdown with discovered printers
function updatePrinterDropdown() {
    const select = document.getElementById('printer-select');
    
    // Clear existing options except the first one (Manual TCP)
    while (select.children.length > 1) {
        select.removeChild(select.lastChild);
    }
    
    // Add discovered printers
    discoveredPrinters.forEach(printer => {
        // Skip the manual TCP entry as it's already in HTML
        if (printer.type === 'manual') {
            return;
        }
        
        const option = document.createElement('option');
        option.value = printer.uri;
        option.textContent = printer.name;
        
        // Add visual indicators for printer types
        if (printer.type === 'local') {
            option.textContent += ' 🖨️';
        } else if (printer.type === 'network') {
            option.textContent += ' 🌐';
        }
        
        select.appendChild(option);
    });
    
    // Update button state
    const refreshButton = document.getElementById('refresh-printers');
    refreshButton.disabled = false;
    refreshButton.querySelector('span:last-child').textContent = 'Refresh';
}

// Show printer discovery errors
function showPrinterError(message) {
    const refreshButton = document.getElementById('refresh-printers');
    refreshButton.disabled = false;
    refreshButton.querySelector('span:last-child').textContent = 'Retry';
    
    // You could add a notification here if desired
    console.error('Printer discovery error:', message);
}

// Handle refresh button click
function handleRefreshClick() {
    const refreshButton = document.getElementById('refresh-printers');
    refreshButton.disabled = true;
    refreshButton.querySelector('span:last-child').textContent = 'Scanning...';
    
    fetchPrinters();
}

// Initialize printer discovery when page loads
document.addEventListener('DOMContentLoaded', function() {
    // Set up refresh button
    const refreshButton = document.getElementById('refresh-printers');
    refreshButton.addEventListener('click', handleRefreshClick);
    
    // Initial printer discovery
    fetchPrinters();
});

// Export functions for potential use by other scripts
window.printerDiscovery = {
    fetchPrinters,
    updatePrinterDropdown,
    discoveredPrinters: () => discoveredPrinters
};
