package main

import (
	"log"
	"os"

	"crm-backend-go/internal/handlers"
	"crm-backend-go/internal/rabbitmq"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get configuration from environment variables
	rabbitMQURL := getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	queueName := getEnv("QUEUE_NAME", "support_tickets_queue")
	port := getEnv("PORT", "8080")

	// Initialize RabbitMQ client
	rabbitMQConfig := rabbitmq.Config{
		URL:       rabbitMQURL,
		QueueName: queueName,
	}

	rabbitMQClient, err := rabbitmq.NewClient(rabbitMQConfig)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQClient.Close()

	// Initialize Gin router
	router := gin.Default()

	// Add middleware for CORS (optional)
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Initialize handlers
	ticketHandler := handlers.NewTicketHandler(rabbitMQClient, queueName)

	// Define routes
	router.POST("/ticket", ticketHandler.CreateTicket)

	// Health check endpoint with RabbitMQ connectivity check
	router.GET("/health", func(c *gin.Context) {
		response := gin.H{
			"status":  "ok",
			"message": "CRM Backend Service is running",
		}

		// Check RabbitMQ connectivity
		if err := rabbitMQClient.HealthCheck(); err != nil {
			response["status"] = "degraded"
			response["message"] = "Service is running but RabbitMQ is not healthy"
			response["rabbitmq_error"] = err.Error()
			c.JSON(503, response) // Service Unavailable
			return
		}

		response["rabbitmq"] = "connected"
		c.JSON(200, response)
	})

	// Readiness probe endpoint for Kubernetes
	router.GET("/ready", func(c *gin.Context) {
		if !rabbitMQClient.IsConnected() {
			c.JSON(503, gin.H{
				"ready":   false,
				"message": "Service not ready - RabbitMQ disconnected",
			})
			return
		}

		c.JSON(200, gin.H{
			"ready":   true,
			"message": "Service is ready",
		})
	})

	// Liveness probe endpoint for Kubernetes
	router.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"alive":   true,
			"message": "Service is alive",
		})
	})

	// Start server
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
