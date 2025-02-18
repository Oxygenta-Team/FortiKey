package services

import (
	"context"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type SecretService struct {
	repoManager repository.RepoManager
}

func NewSecretService(repoManager repository.RepoManager) SecretSvc {
	return &SecretService{repoManager: repoManager}
}

func (s *SecretService) InsertSecret(ctx context.Context, secret []*models.Secret) error {
	//TODO implement me
	panic("implement me")
}

func (s *SecretService) GetSecretByID(ctx context.Context, id uint64) (*models.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SecretService) GetSecretByKey(ctx context.Context, key string) (*models.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SecretService) DeleteSecret(ctx context.Context, ids []uint64) error {
	//TODO implement me
	panic("implement me")
}
