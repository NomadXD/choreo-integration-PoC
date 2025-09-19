package handlers

import (
	"net/http"

	"crm-backend-go/internal/models"
	"crm-backend-go/internal/rabbitmq"

	"github.com/gin-gonic/gin"
)

// TicketHandler handles ticket-related HTTP requests
type TicketHandler struct {
	rabbitMQClient *rabbitmq.Client
	queueName      string
}

// NewTicketHandler creates a new ticket handler
func NewTicketHandler(rabbitMQClient *rabbitmq.Client, queueName string) *TicketHandler {
	return &TicketHandler{
		rabbitMQClient: rabbitMQClient,
		queueName:      queueName,
	}
}

// CreateTicket handles POST /ticket requests
func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var ticket models.Ticket

	// Bind JSON payload to ticket struct
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, models.TicketResponse{
			Success: false,
			Message: "Invalid request payload: " + err.Error(),
		})
		return
	}

	// Publish ticket to RabbitMQ
	if err := h.rabbitMQClient.PublishTicket(h.queueName, ticket); err != nil {
		c.JSON(http.StatusInternalServerError, models.TicketResponse{
			Success: false,
			Message: "Failed to process ticket: " + err.Error(),
		})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, models.TicketResponse{
		Success:  true,
		Message:  "Ticket submitted successfully",
		TicketID: ticket.TicketID,
	})
}
