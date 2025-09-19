package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Client represents a RabbitMQ client
type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Config holds RabbitMQ connection configuration
type Config struct {
	URL       string
	QueueName string
}

// NewClient creates a new RabbitMQ client
func NewClient(config Config) (*Client, error) {
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	// Declare the queue
	// _, err = ch.QueueDeclare(
	// 	config.QueueName, // name
	// 	true,             // durable
	// 	false,            // delete when unused
	// 	false,            // exclusive
	// 	false,            // no-wait
	// 	nil,              // arguments
	// )
	// if err != nil {
	// 	ch.Close()
	// 	conn.Close()
	// 	return nil, fmt.Errorf("failed to declare queue: %w", err)
	// }

	return &Client{
		conn:    conn,
		channel: ch,
	}, nil
}

// PublishTicket publishes a ticket to the RabbitMQ queue
func (c *Client) PublishTicket(queueName string, ticket interface{}) error {
	body, err := json.Marshal(ticket)
	if err != nil {
		return fmt.Errorf("failed to marshal ticket: %w", err)
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
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Published ticket to queue %s: %s", queueName, string(body))
	return nil
}

// Close closes the RabbitMQ connection and channel
func (c *Client) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
