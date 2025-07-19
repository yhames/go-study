package kafka

import (
	"chat-ws-control/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Kafka struct {
	config   *config.Config
	consumer *kafka.Consumer
}

func NewKafka(config *config.Config) (*Kafka, error) {
	k := &Kafka{config: config}
	var err error
	// TODO: Change consumer
	k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Url,
		"group.id":          config.Kafka.GroupId,
		"auto.offset.reset": "latest",
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (k *Kafka) Subscribe(topic string) error {
	return k.consumer.Subscribe(topic, nil)
}

func (k *Kafka) Poll(timeoutMs int) kafka.Event {
	return k.consumer.Poll(timeoutMs)
}
