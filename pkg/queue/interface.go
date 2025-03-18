package queue

import (
	"context"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type Producer interface {
	ProduceMessages(ctx context.Context, rawMessages []*models.KafkaMessage) error
	Close() error
}
