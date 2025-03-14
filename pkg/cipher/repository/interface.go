package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type RepoManager interface {
	NewSecretRepo(db sqlx.ExtContext) SecretRepo
}

type SecretRepo interface {
	InsertSecret(ctx context.Context, secret []*models.Secret) error
	// TODO make `FetchSecret` with filters
	GetSecretByID(ctx context.Context, id uint64) (*models.Secret, error)
	GetSecretByKey(ctx context.Context, key string) (*models.Secret, error)
	DeleteSecret(ctx context.Context, ids []uint64) error
}
