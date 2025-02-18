package services

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
)

func NewServices(repoManager repository.RepoManager, storage *postgres.Storage) *Services {
	return &Services{
		SecretSvc: NewSecretService(repoManager, storage),
	}
}
