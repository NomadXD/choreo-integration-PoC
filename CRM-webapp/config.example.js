// Example configuration for different environments
// Copy this file and rename to config.js, then modify as needed

const CONFIG = {
  // Development environment
  BACKEND_URL: "http://localhost:3000/api/tickets",

  // Production environment (uncomment and modify for production)
  // BACKEND_URL: 'https://api.yourcompany.com/tickets',

  // Staging environment (uncomment and modify for staging)
  // BACKEND_URL: 'https://staging-api.yourcompany.com/tickets',

  REQUEST_TIMEOUT: 10000,

  API_HEADERS: {
    "Content-Type": "application/json",

    // Add authentication headers if your API requires them
    // 'Authorization': 'Bearer your-api-token-here',
    // 'X-API-Key': 'your-api-key-here',

    // Add any custom headers your backend expects
    // 'X-Client-Version': '1.0.0',
    // 'X-Requested-With': 'SupportTicketApp'
  },
};

// For Node.js environments (if using this config server-side)
if (typeof module !== "undefined" && module.exports) {
  module.exports = CONFIG;
}
