// Package pubsub contains several pubsub methods which implement
// common interface.
package pubsub

// PubSub is pubsub interface.
//
// Subscribe function should return Channel interface
// but Go doesn't allow interface method to return another
// interface. So, you have to convert it first before using it.
//
// See usage example in example folder.
type PubSub interface {
	// Publish message to specific topic/channel.
	// Data will be encoded first before publishing.
	Publish(topic string, data interface{}) error
	// Subscribe to specific topic/channel.
	Subscribe(topic string) (interface{}, error)
	// Close pubsub client connection.
	Close() error
}

// Channel is channel interface.
//
// See usage example in example folder.
type Channel interface {
	// Read and process incoming message.
	Read(model interface{}) (<-chan interface{}, <-chan error)
	// Close subscription.
	Close() error
}
