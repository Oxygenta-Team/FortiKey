package services

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
)

func NewServices(repoManager repository.RepoManager, storage *postgres.Storage, logger *logging.Logger) *Services {
	return &Services{
		SecretSvc: NewSecretService(repoManager, storage, logger.WithField("component", "SecretService")),
	}
}
