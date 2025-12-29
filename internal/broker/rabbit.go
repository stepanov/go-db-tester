package broker

import "errors"

// NewRabbit is a stub for RabbitMQ broker creation
func NewRabbit(url string) (Broker, error) {
	return nil, errors.New("rabbit broker not implemented")
}
