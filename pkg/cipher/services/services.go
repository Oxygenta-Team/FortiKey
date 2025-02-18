package services

import "github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"

func NewServices(repoManager repository.RepoManager) *Services {
	return &Services{
		SecretSvc: NewSecretService(repoManager),
	}
}
