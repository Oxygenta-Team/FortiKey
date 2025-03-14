package kafka

import (
	"context"
	"encoding/json"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	kafkaWriter *kafka.Writer
}

func NewProducer(config *Config) *Producer {
	return &Producer{kafkaWriter: &kafka.Writer{
		Addr:  kafka.TCP(config.Brokers...),
		Topic: config.Producer.Topic,

		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Compression:  kafka.Snappy,
	}}
}

func (producer *Producer) ProduceMessages(ctx context.Context, rawMessages []*models.KafkaMessage) error {
	messages := make([]kafka.Message, len(rawMessages))
	for i, value := range rawMessages {
		m, err := json.Marshal(value)
		if err != nil {
			return err
		}
		messages[i] = kafka.Message{
			Value: m,
		}
	}

	return producer.kafkaWriter.WriteMessages(ctx, messages...)
}

func (producer *Producer) Close() error {
	return producer.kafkaWriter.Close()
}
