package services

import (
	"encoding/json"

	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka"
)

func (s *Services) StartConsumer(l *logging.Logger, config *kafka.Config) {
	topicsToStartFactConsumerFunctions := kafka.TopicToConsumer{
		"cipher-updates": s.StartUserManagementConsumer,
	}

	kafka.StartConsuming(l, config, topicsToStartFactConsumerFunctions)
}

func (s *Services) StartUserManagementConsumer(l *logging.Logger, conf *kafka.Config, topic string) {
	kafka.ConsumeMessages(l, conf, topic, func(message *kafka.Message) error {
		//ctx := context.Background()
		var obj *models.KafkaMessage
		l.Warnf("take a message from user-management, msg:%+v", message.Message)

		err := json.Unmarshal(message.Message.Value, &obj)
		if err != nil {
			return err
		}

		switch obj.ObjectType {
		case models.SecretObjectType:
			// Add user
			var secret models.Secret
			err := json.Unmarshal(*obj.Object, &secret)
			if err != nil {
				return err
			}
			// Do some
		}

		return nil
	})
}
