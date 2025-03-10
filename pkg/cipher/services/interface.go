package services

import (
	"context"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type Services struct {
	SecretSvc SecretSvc
}

type SecretSvc interface {
	InsertSecret(ctx context.Context, secret []*models.Secret) error
	// TODO make `FetchSecret` with filters
	GetSecretByID(ctx context.Context, id uint64) (*models.Secret, error)
	GetSecretByKey(ctx context.Context, key string) (*models.Secret, error)
	DeleteSecret(ctx context.Context, ids []uint64) error
}
