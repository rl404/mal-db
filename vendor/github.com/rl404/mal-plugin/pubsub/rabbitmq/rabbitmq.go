// Package rabbitmq is a wrapper of the original "github.com/streadway/amqp" library.
//
// Only contains basic publish, subscribe, and close methods.
// Data will be encoded to JSON before publishing the message.
package rabbitmq

import (
	"encoding/json"

	"github.com/rl404/mal-plugin/pubsub"
	"github.com/streadway/amqp"
)

// Rabbitmq client implements PubSub interface.
var _ pubsub.PubSub = &Client{}

// Rabbitmq channels implements Channel interface.
var _ pubsub.Channel = &Channel{}

// Client is rabbitmq pubsub client.
type Client struct {
	client *amqp.Connection
}

// Channel is rabbitmq subscription channel.
type Channel struct {
	channel  *amqp.Channel
	messages <-chan amqp.Delivery
}

// New to create new rabbitmq pubsub client.
func New(url string) (*Client, error) {
	c, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client: c}, nil
}

// Publish to publish message.
func (c *Client) Publish(queue string, data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ch, err := c.client.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        j,
	})
	if err != nil {
		return err
	}

	return ch.Close()
}

// Subscribe to subscribe queue.
func (c *Client) Subscribe(queue string) (interface{}, error) {
	ch, err := c.client.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &Channel{
		channel:  ch,
		messages: msgs,
	}, nil
}

// Close to close pubsub connection.
func (c *Client) Close() error {
	return c.client.Close()
}

// Read to read incoming message.
func (c *Channel) Read(model interface{}) (<-chan interface{}, <-chan error) {
	msgChan, errChan := make(chan interface{}), make(chan error)
	go func() {
		for msg := range c.messages {
			if err := json.Unmarshal(msg.Body, &model); err != nil {
				errChan <- err
			} else {
				msgChan <- model
			}
		}
	}()
	return (<-chan interface{})(msgChan), (<-chan error)(errChan)
}

// Close to close subscription.
func (c *Channel) Close() error {
	return c.channel.Close()
}
