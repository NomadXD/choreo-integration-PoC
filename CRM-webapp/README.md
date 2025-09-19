# Support Ticket System

A simple web application for creating support tickets by calling a remote backend API endpoint.

## Features

- Clean, responsive web interface
- Form validation with real-time feedback
- Configurable backend URL
- Loading states and error handling
- Success/failure feedback messages
- Mobile-friendly design

## Setup

1. **Clone or download** this project to your local machine

2. **Configure the backend URL**:

   - Open `config.js`
   - Replace `'https://your-backend-api.com/api/tickets'` with your actual backend API endpoint
   - Add any required headers (like authentication tokens) in the `API_HEADERS` section

3. **Serve the files**:
   - You can open `index.html` directly in a browser for testing
   - For production, serve through a web server (Apache, Nginx, etc.)
   - For local development, you can use a simple HTTP server:

     ```bash
     # Using Python
     python -m http.server 8000

     # Using Node.js (if you have http-server installed)
     npx http-server

     # Using PHP
     php -S localhost:8000
     ```

## Configuration

### Backend URL Configuration

Edit `config.js` to configure your backend:

```javascript
const CONFIG = {
  // Your backend API endpoint
  BACKEND_URL: "https://your-backend-api.com/api/tickets",

  // Request timeout (in milliseconds)
  REQUEST_TIMEOUT: 10000,

  // API headers
  API_HEADERS: {
    "Content-Type": "application/json",
    // Add authentication headers if needed
    // 'Authorization': 'Bearer your-token-here'
  },
};
```

### Expected Backend API

Your backend should accept POST requests with the following JSON payload:

```json
{
  "ticket_id": "INC-69",
  "customer_name": "Jane Doe",
  "customer_email": "lahiru97udayanga@gmail.com",
  "issue": "Cannot reset password."
}
```

The API should return a JSON response on success.

## Form Fields

- **Ticket ID**: Required, format ABC-123 (3 letters, dash, numbers)
- **Customer Name**: Required, minimum 2 characters
- **Customer Email**: Required, valid email format
- **Issue Description**: Required, minimum 10 characters

## Form Validation

The form includes both client-side validation and user-friendly error messages:

- Real-time validation on field blur
- Format validation for ticket ID and email
- Minimum length requirements
- Visual feedback with error styling

## Error Handling

The application handles various error scenarios:

- Network connectivity issues
- Request timeouts
- Server errors (4xx, 5xx responses)
- Invalid backend URL configuration
- Form validation errors

## Browser Compatibility

This application uses modern web standards and is compatible with:

- Chrome/Chromium (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Files Structure

```
├── index.html          # Main HTML page
├── styles.css          # CSS styling
├── script.js           # JavaScript functionality
├── config.js           # Configuration file
└── README.md           # This file
```

## Customization

### Styling

Modify `styles.css` to change the appearance:

- Colors and gradients
- Typography
- Layout and spacing
- Responsive breakpoints

### Functionality

Extend `script.js` to add features like:

- Additional form fields
- Different validation rules
- Custom error handling
- Analytics tracking

### Configuration

Update `config.js` for:

- Different backend endpoints
- Custom headers
- Timeout settings
- API authentication

## Security Considerations

- All form inputs are validated both client-side and should be validated server-side
- Use HTTPS for production deployments
- Implement proper CORS policies on your backend
- Consider implementing rate limiting
- Sanitize all user inputs on the backend

## Troubleshooting

### Common Issues

1. **"Unable to connect to the server"**

   - Check that the backend URL in `config.js` is correct
   - Verify the backend service is running
   - Check for CORS issues

2. **"Request timed out"**

   - Increase the `REQUEST_TIMEOUT` value in `config.js`
   - Check network connectivity
   - Verify backend response time

3. **Form validation errors**
   - Ensure all required fields are filled
   - Check format requirements (especially Ticket ID and email)
   - Verify minimum length requirements

## Development

To modify or extend this application:

1. Make changes to the relevant files
2. Test in multiple browsers
3. Validate with different screen sizes
4. Test error scenarios
5. Update this README if adding new features

## License

This project is provided as-is for educational and development purposes.
