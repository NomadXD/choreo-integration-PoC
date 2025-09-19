// Support Ticket System JavaScript

class TicketSystem {
  constructor() {
    this.form = document.getElementById("ticketForm");
    this.submitBtn = document.getElementById("submitBtn");
    this.btnText = document.querySelector(".btn-text");
    this.btnSpinner = document.querySelector(".btn-spinner");
    this.messageDiv = document.getElementById("message");

    this.init();
  }

  init() {
    this.form.addEventListener("submit", this.handleSubmit.bind(this));
    this.addInputValidation();
  }

  addInputValidation() {
    const inputs = this.form.querySelectorAll("input, textarea");
    inputs.forEach((input) => {
      input.addEventListener("blur", this.validateField.bind(this));
      input.addEventListener("input", this.clearFieldError.bind(this));
    });
  }

  validateField(event) {
    const field = event.target;
    const value = field.value.trim();

    // Remove any existing error styling
    field.classList.remove("error");

    // Validate based on field type
    switch (field.name) {
      case "ticket_id":
        if (!value) {
          this.setFieldError(field, "Ticket ID is required");
        } else if (!/^[A-Z]{3}-\d+$/.test(value)) {
          this.setFieldError(field, "Ticket ID should be in format: ABC-123");
        }
        break;

      case "customer_name":
        if (!value) {
          this.setFieldError(field, "Customer name is required");
        } else if (value.length < 2) {
          this.setFieldError(
            field,
            "Customer name must be at least 2 characters"
          );
        }
        break;

      case "customer_email":
        if (!value) {
          this.setFieldError(field, "Email is required");
        } else if (!this.isValidEmail(value)) {
          this.setFieldError(field, "Please enter a valid email address");
        }
        break;

      case "issue":
        if (!value) {
          this.setFieldError(field, "Issue description is required");
        } else if (value.length < 10) {
          this.setFieldError(
            field,
            "Please provide a more detailed description (at least 10 characters)"
          );
        }
        break;
    }
  }

  setFieldError(field, message) {
    field.classList.add("error");

    // Remove any existing error message
    const existingError = field.parentNode.querySelector(".field-error");
    if (existingError) {
      existingError.remove();
    }

    // Add new error message
    const errorDiv = document.createElement("div");
    errorDiv.className = "field-error";
    errorDiv.textContent = message;
    errorDiv.style.color = "#dc3545";
    errorDiv.style.fontSize = "0.85rem";
    errorDiv.style.marginTop = "4px";

    field.parentNode.appendChild(errorDiv);
  }

  clearFieldError(event) {
    const field = event.target;
    field.classList.remove("error");

    const errorDiv = field.parentNode.querySelector(".field-error");
    if (errorDiv) {
      errorDiv.remove();
    }
  }

  isValidEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }

  validateForm() {
    const inputs = this.form.querySelectorAll("input, textarea");
    let isValid = true;

    inputs.forEach((input) => {
      this.validateField({ target: input });
      if (input.classList.contains("error") || !input.value.trim()) {
        isValid = false;
      }
    });

    return isValid;
  }

  async handleSubmit(event) {
    event.preventDefault();

    // Validate form
    if (!this.validateForm()) {
      this.showMessage(
        "Please fix the errors in the form before submitting.",
        "error"
      );
      return;
    }

    // Get form data
    const formData = new FormData(this.form);
    const ticketData = {
      ticket_id: formData.get("ticket_id").trim(),
      customer_name: formData.get("customer_name").trim(),
      customer_email: formData.get("customer_email").trim(),
      issue: formData.get("issue").trim(),
    };

    try {
      await this.createTicket(ticketData);
    } catch (error) {
      console.error("Error creating ticket:", error);
      this.showMessage("Failed to create ticket. Please try again.", "error");
    }
  }

  async createTicket(ticketData) {
    // Show loading state
    this.setLoadingState(true);
    this.hideMessage();

    try {
      // Check if CONFIG is available
      if (typeof CONFIG === "undefined") {
        throw new Error("Configuration not loaded");
      }

      // Create abort controller for timeout
      const controller = new AbortController();
      const timeoutId = setTimeout(
        () => controller.abort(),
        CONFIG.REQUEST_TIMEOUT
      );

      const response = await fetch(CONFIG.BACKEND_URL, {
        method: "POST",
        headers: CONFIG.API_HEADERS,
        body: JSON.stringify(ticketData),
        signal: controller.signal,
      });

      clearTimeout(timeoutId);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();

      // Success
      this.showMessage(
        `Ticket created successfully! Ticket ID: ${ticketData.ticket_id}`,
        "success"
      );
      this.form.reset();
    } catch (error) {
      if (error.name === "AbortError") {
        this.showMessage(
          "Request timed out. Please check your connection and try again.",
          "error"
        );
      } else if (error.message.includes("Failed to fetch")) {
        this.showMessage(
          "Unable to connect to the server. Please check the backend URL in config.js and try again.",
          "error"
        );
      } else {
        this.showMessage(`Error: ${error.message}`, "error");
      }
      throw error;
    } finally {
      this.setLoadingState(false);
    }
  }

  setLoadingState(isLoading) {
    this.submitBtn.disabled = isLoading;

    if (isLoading) {
      this.btnText.style.display = "none";
      this.btnSpinner.style.display = "inline-flex";
    } else {
      this.btnText.style.display = "inline";
      this.btnSpinner.style.display = "none";
    }
  }

  showMessage(text, type) {
    this.messageDiv.textContent = text;
    this.messageDiv.className = `message ${type}`;
    this.messageDiv.style.display = "block";

    // Auto-hide success messages after 5 seconds
    if (type === "success") {
      setTimeout(() => {
        this.hideMessage();
      }, 5000);
    }
  }

  hideMessage() {
    this.messageDiv.style.display = "none";
  }
}

// Initialize the ticket system when the DOM is loaded
document.addEventListener("DOMContentLoaded", () => {
  new TicketSystem();
});

// Add some additional CSS for field errors
const style = document.createElement("style");
style.textContent = `
    input.error, textarea.error {
        border-color: #dc3545 !important;
        background-color: #fff5f5 !important;
    }
    
    .field-error {
        animation: slideIn 0.3s ease-out;
    }
    
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(-10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
`;
document.head.appendChild(style);
