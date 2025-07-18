// static/js/settings.js

// Settings panel toggle functionality

document.addEventListener('DOMContentLoaded', function() {
    const settingsToggle = document.getElementById('settings-toggle');
    const settingsPanel = document.getElementById('settings-panel');
    const settingsChevron = document.getElementById('settings-chevron');

    let isExpanded = false;

    // Toggle settings panel visibility
    function toggleSettings() {
        isExpanded = !isExpanded;

        if (isExpanded) {
            settingsPanel.classList.remove('is-hidden');
            settingsChevron.textContent = '▲';
            settingsToggle.classList.add('is-active');
        } else {
            settingsPanel.classList.add('is-hidden');
            settingsChevron.textContent = '▼';
            settingsToggle.classList.remove('is-active');
        }
    }

    // Add click event listener to toggle button
    settingsToggle.addEventListener('click', function(e) {
        e.preventDefault();
        toggleSettings();
    });

    // Optional: Close settings panel when clicking outside
    document.addEventListener('click', function(e) {
        const settingsBox = settingsToggle.closest('.box');
        if (isExpanded && !settingsBox.contains(e.target)) {
            toggleSettings();
        }
    });

    // Optional: Keyboard accessibility (Enter/Space to toggle)
    settingsToggle.addEventListener('keydown', function(e) {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            toggleSettings();
        }
    });
});

// Export for potential use by other scripts
window.settingsPanel = {
    toggle: function() {
        document.getElementById('settings-toggle').click();
    }
};