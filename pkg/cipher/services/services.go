package services

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/redis"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue"
)

func NewServices(repoManager repository.RepoManager, producer queue.Producer, storage *postgres.Storage, rcl redis.Cacher, logger *logging.Logger) *Services {
	return &Services{
		SecretSvc: NewSecretService(repoManager, producer, storage, rcl, logger.WithField("component", "SecretService")),
	}
}
