package services

import (
	"context"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

type SecretService struct {
	repoManager repository.RepoManager
	db          *postgres.Storage
}

func NewSecretService(repoManager repository.RepoManager, db *postgres.Storage) SecretSvc {
	return &SecretService{repoManager: repoManager, db: db}
}

func (s *SecretService) CreateSecret(ctx context.Context, secrets []*models.Secret) error {
	secRepo := s.repoManager.NewSecretRepo(s.db)

	for i, secret := range secrets {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(secret.Value), bcrypt.MaxCost)
		if err != nil {
			return err
		}
		secrets[i].Hash = hashedPassword
	}

	err := secRepo.InsertSecret(ctx, secrets)
	if err != nil {
		return err
	}

	return err
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
