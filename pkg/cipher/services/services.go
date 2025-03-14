package services

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka"
)

func NewServices(repoManager repository.RepoManager, producer *kafka.Producer, storage *postgres.Storage, logger *logging.Logger) *Services {
	return &Services{
		SecretSvc: NewSecretService(repoManager, producer, storage, logger.WithField("component", "SecretService")),
	}
}
