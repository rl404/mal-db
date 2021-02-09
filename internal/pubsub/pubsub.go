package pubsub

import (
	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-plugin/pubsub/nsq"
	"github.com/rl404/mal-plugin/pubsub/rabbitmq"
	"github.com/rl404/mal-plugin/pubsub/redis"
)

// PubSub is pubsub interface.
type PubSub interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string) (interface{}, error)
	Close() error
}

// Channel is channel interface.
type Channel interface {
	Read(model interface{}) (<-chan interface{}, <-chan error)
	Close() error
}

// Message is pubsub message model.
type Message struct {
	Type string
	ID   int
}

// New to create new pubsub client.
func New(l logger.Logger, pubsubType string, address string, user string, password string) (PubSub, error) {
	switch pubsubType {
	case config.Redis:
		l.Debug("using redis pubsub")
		return redis.New(address, password)
	case config.NSQ:
		l.Debug("using NSQ pubsub")
		return nsq.New(address)
	case config.RabbitMQ:
		l.Debug("using rabbitmq pubsub")
		return rabbitmq.New(address)
	default:
		return nil, errors.ErrRequiredPubsub
	}
}
