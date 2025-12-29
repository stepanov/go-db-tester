package broker

import (
	"errors"
)

// NewNats is a stub that would create a NATS-based Broker implementation
func NewNats(url string) (Broker, error) {
	// In a real implementation, construct nats.go client
	return nil, errors.New("nats broker not implemented")
}
