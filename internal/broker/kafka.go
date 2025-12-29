package broker

import (
	"errors"
)

// NewKafka is a stub that would create a Kafka-based Broker implementation
func NewKafka(brokers []string) (Broker, error) {
	// In a real implementation, construct sarama or confluent-kafka client
	return nil, errors.New("kafka broker not implemented")
}
