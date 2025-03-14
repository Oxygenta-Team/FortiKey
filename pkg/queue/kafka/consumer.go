package kafka

import (
	"context"
	"errors"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

type Message struct {
	*kafka.Message
}

type MessageHandler func(*Message) error
type TopicToConsumer map[string]func(l *logging.Logger, conf *Config, topic string)

func StartConsuming(l *logging.Logger, conf *Config, topicToConsumer TopicToConsumer) {
	for topic, startFactConsumerFunction := range topicToConsumer {
		for range conf.Consumer.ConsumersAmount {
			go startFactConsumerFunction(l, conf, topic)
		}
	}
}

func ConsumeMessages(l *logging.Logger, conf *Config, topic string, msgHandler MessageHandler) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.Brokers,
		GroupID:  conf.Consumer.Group,
		MaxWait:  500 * time.Millisecond,
		MinBytes: 1,
	})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		if err := r.Close(); err != nil {
			l.Errorf("could not close kafka reader. err: %s)", err)
		}
	}()

	for {
		rawMsg, err := r.ReadMessage(context.Background())
		if err != nil {
			if errors.Is(err, io.EOF) {
				_ = r.CommitMessages(context.Background(), rawMsg)
			}
			continue
		}
		if rawMsg.Value != nil {
			l.Errorf("could not handle kafka message. err: %s)", err)
			err = msgHandler(&Message{&rawMsg})
		}
	}
}
