// Configuration file for the Support Ticket System
// Modify this file to change the backend URL

const CONFIG = {
  // Backend API endpoint for creating support tickets
  BACKEND_URL:
    "https://56792605-ed4d-4cc1-95a1-fbad23a01bff-dev.e1-us-east-azure.preview-dv.choreoapis.dev/pldr/crm-backend-go/v1.0/ticket",

  // Request timeout in milliseconds
  REQUEST_TIMEOUT: 30000,

  // API headers (if needed)
  API_HEADERS: {
    "Content-Type": "application/json",
    // Add any other headers like authentication tokens here
    // 'Authorization': 'Bearer your-token-here'
  },
};

// Export for use in other files
if (typeof module !== "undefined" && module.exports) {
  module.exports = CONFIG;
}
