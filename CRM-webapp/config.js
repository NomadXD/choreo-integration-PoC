// Configuration file for the Support Ticket System
// Modify this file to change the backend URL

const CONFIG = {
  // Backend API endpoint for creating support tickets
  BACKEND_URL:
    "https://b48cc93e-fa33-4420-a155-bc653b4d46be-my-env.e1-us-east-azure.choreoapis.dev/aaa-poc/crm-backend/v1.0/ticket",

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
