package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository"
)

type RepoManager struct {
}

func NewRepoManager() *RepoManager {
	return &RepoManager{}
}

func (r *RepoManager) NewSecretRepo(db sqlx.ExtContext) repository.SecretRepo {
	return NewSecretRepo(db)
}
