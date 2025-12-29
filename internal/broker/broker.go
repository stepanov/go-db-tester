package broker

import "context"

// Broker is a minimal interface for message brokers used by the app
type Broker interface {
	Publish(ctx context.Context, topic string, key []byte, msg []byte) error
	Subscribe(ctx context.Context, topic string, handler func([]byte) error) error
	Close() error
}

// Provider allows us to hold multiple broker implementations
type Provider struct {
	Kafka  Broker
	Nats   Broker
	Rabbit Broker
}
