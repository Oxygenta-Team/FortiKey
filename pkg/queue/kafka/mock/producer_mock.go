package kafka_mock

import (
	"context"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue"
)

type ProducerMock struct {
	errorToReturn error
}

func NewProducerMock(errorToReturn error) queue.Producer {
	return &ProducerMock{errorToReturn: errorToReturn}
}

func (p *ProducerMock) ProduceMessages(ctx context.Context, rawMessages []*models.KafkaMessage) error {
	return p.errorToReturn
}

func (p *ProducerMock) Close() error {
	return p.errorToReturn
}
