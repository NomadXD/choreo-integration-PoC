package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Client represents a RabbitMQ client
type Client struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	config    Config
	mutex     sync.RWMutex
	connected bool
	closeChan chan *amqp.Error
	reconnect chan bool
	done      chan bool
}

// Config holds RabbitMQ connection configuration
type Config struct {
	URL       string
	QueueName string
}

// NewClient creates a new RabbitMQ client with auto-reconnection
func NewClient(config Config) (*Client, error) {
	client := &Client{
		config:    config,
		reconnect: make(chan bool),
		done:      make(chan bool),
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}

	// Start reconnection goroutine
	go client.handleReconnect()

	return client, nil
}

// connect establishes connection to RabbitMQ
func (c *Client) connect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var err error
	c.conn, err = amqp.Dial(c.config.URL)
	if err != nil {
		c.connected = false
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		c.conn.Close()
		c.connected = false
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	// Set up connection close handler
	c.closeChan = make(chan *amqp.Error)
	c.conn.NotifyClose(c.closeChan)

	// Declare the queue if needed
	// if c.config.QueueName != "" {
	// 	_, err = c.channel.QueueDeclare(
	// 		c.config.QueueName, // name
	// 		true,               // durable
	// 		false,              // delete when unused
	// 		false,              // exclusive
	// 		false,              // no-wait
	// 		nil,                // arguments
	// 	)
	// 	if err != nil {
	// 		c.channel.Close()
	// 		c.conn.Close()
	// 		c.connected = false
	// 		return fmt.Errorf("failed to declare queue: %w", err)
	// 	}
	// }

	c.connected = true
	log.Println("Connected to RabbitMQ")
	return nil
}

// handleReconnect handles automatic reconnection
func (c *Client) handleReconnect() {
	for {
		select {
		case err := <-c.closeChan:
			if err != nil {
				log.Printf("RabbitMQ connection closed: %v", err)
				c.connected = false

				// Attempt to reconnect
				for {
					log.Println("Attempting to reconnect to RabbitMQ...")
					if err := c.connect(); err != nil {
						log.Printf("Failed to reconnect: %v", err)
						time.Sleep(5 * time.Second)
						continue
					}
					log.Println("Successfully reconnected to RabbitMQ")
					break
				}
			}
		case <-c.done:
			return
		}
	}
}

// IsConnected returns the current connection status
func (c *Client) IsConnected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.connected && c.conn != nil && !c.conn.IsClosed()
}

// HealthCheck performs a health check by attempting to inspect the queue
func (c *Client) HealthCheck() error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if !c.connected || c.conn == nil || c.conn.IsClosed() {
		return fmt.Errorf("not connected to RabbitMQ")
	}

	if c.channel == nil {
		return fmt.Errorf("no active channel")
	}

	// Try to inspect a queue to verify the connection is working
	if c.config.QueueName != "" {
		_, err := c.channel.QueueInspect(c.config.QueueName)
		if err != nil {
			return fmt.Errorf("failed to inspect queue: %w", err)
		}
	}

	return nil
}

// PublishTicket publishes a ticket to the RabbitMQ queue with retry logic
func (c *Client) PublishTicket(queueName string, ticket interface{}) error {
	body, err := json.Marshal(ticket)
	if err != nil {
		return fmt.Errorf("failed to marshal ticket: %w", err)
	}

	// Retry publishing up to 3 times
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		c.mutex.RLock()
		if !c.connected || c.channel == nil {
			c.mutex.RUnlock()
			if i == maxRetries-1 {
				return fmt.Errorf("not connected to RabbitMQ after %d attempts", maxRetries)
			}
			time.Sleep(1 * time.Second)
			continue
		}

		err = c.channel.Publish(
			"",        // exchange
			queueName, // routing key (queue name)
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		c.mutex.RUnlock()

		if err == nil {
			log.Printf("Published ticket to queue %s: %s", queueName, string(body))
			return nil
		}

		log.Printf("Failed to publish message (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(1 * time.Second)
		}
	}

	return fmt.Errorf("failed to publish message after %d attempts: %w", maxRetries, err)
}

// Close closes the RabbitMQ connection and channel
func (c *Client) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Stop the reconnect goroutine
	close(c.done)

	c.connected = false

	if c.channel != nil {
		c.channel.Close()
		c.channel = nil
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}

	log.Println("RabbitMQ client closed")
}
