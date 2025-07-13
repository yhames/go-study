package kafka

import (
	"chat-ws/app/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Kafka struct {
	config   *config.Config
	producer *kafka.Producer
}

func NewKafka(config *config.Config) (*Kafka, error) {
	k := &Kafka{config: config}
	var err error
	k.producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Url,
		"client.id":         config.Kafka.ClientId,
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}
	return k, nil
}
