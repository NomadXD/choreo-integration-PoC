package models

// Ticket represents a customer support ticket
type Ticket struct {
	TicketID      string `json:"ticket_id" binding:"required"`
	CustomerName  string `json:"customer_name" binding:"required"`
	CustomerEmail string `json:"customer_email" binding:"required,email"`
	Issue         string `json:"issue" binding:"required"`
}

// TicketResponse represents the response sent back to the client
type TicketResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	TicketID string `json:"ticketId,omitempty"`
}
