package services

import (
	"context"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/crypt"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type SecretService struct {
	repoManager repository.RepoManager
	db          *postgres.Storage
	logger      *logging.Logger
}

func NewSecretService(repoManager repository.RepoManager, db *postgres.Storage, logger *logging.Logger) SecretSvc {
	return &SecretService{repoManager: repoManager, db: db, logger: logger}
}

func (s *SecretService) CreateSecret(ctx context.Context, secrets []*models.Secret) error {
	var err error
	defer func() {
		if err != nil {
			s.logger.Errorf("error during creating secret, err:%s, secrets: %v", err, secrets)
		}
	}()
	secRepo := s.repoManager.NewSecretRepo(s.db)
	for _, secret := range secrets {
		if err = crypt.BCryptSecret(secret); err != nil {
			return err
		}
	}

	err = secRepo.InsertSecret(ctx, secrets)
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

func (s *SecretService) CompareSecret(ctx context.Context, keyValue *models.KeyValue) (bool, error) {
	var err error
	defer func() {
		if err != nil {
			s.logger.Errorf("error during creating secret, err:%s, keyValue: %v", err, keyValue)
		}
	}()

	secret, err := s.repoManager.NewSecretRepo(s.db).GetSecretByKey(ctx, keyValue.Key)
	if err != nil {
		return false, err
	}

	switch secret.Method {
	case models.BCRYPT:
		ok := crypt.BCryptCompare(secret.Hash, keyValue.Value)
		if !ok {
			return false, crypt.ErrBcryptCompare
		} else {
			return ok, nil
		}
	}

	return false, ErrInternal
}

func (s *SecretService) DeleteSecret(ctx context.Context, ids []uint64) error {
	//TODO implement me
	panic("implement me")
}
