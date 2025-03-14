package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"

	sq "github.com/Masterminds/squirrel"
)

type SecretRepository struct {
	db sqlx.ExtContext
}

func NewSecretRepo(db sqlx.ExtContext) *SecretRepository {
	return &SecretRepository{db: db}
}

func (s *SecretRepository) InsertSecret(ctx context.Context, secrets []*models.Secret) error {
	insert := sq.
		Insert("secrets").
		Columns(`
			user_id,			
			key,
			method,
			hash
		`)
	for _, secret := range secrets {
		insert = insert.Values(
			secret.UserID,
			secret.Key,
			secret.Method,
			secret.Hash,
		)
	}
	q, args, err := insert.Suffix("RETURNING id").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	rows, err := s.db.QueryxContext(ctx, q, args...)
	if err != nil {
		return err
	}
	for i := 0; rows.Next(); i++ {
		if err := rows.Scan(&secrets[i].ID); err != nil {
			return err
		}
	}

	return err
}

func (s *SecretRepository) GetSecretByID(ctx context.Context, id uint64) (*models.Secret, error) {
	q, args, err := sq.
		Select(`
			s.id,
			s.user_id,
			s.key,
			s.method,
			s.hash
		`).
		From("secrets s").
		Where(sq.Eq{
			"s.id":         id,
			"s.is_deleted": false,
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var secret models.Secret
	err = sqlx.GetContext(ctx, s.db, &secret, q, args...)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func (s *SecretRepository) GetSecretByKey(ctx context.Context, key string) (*models.Secret, error) {
	q, args, err := sq.Select(`
			s.id, 
			s.user_id,
			s.key,
			s.method,
			s.hash
		`).
		From("secrets s").
		Where(sq.Eq{
			"s.key":        key,
			"s.is_deleted": false,
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var secret models.Secret
	err = sqlx.GetContext(ctx, s.db, &secret, q, args...)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func (s *SecretRepository) DeleteSecret(ctx context.Context, ids []uint64) error {
	q, args, err := sq.Update("secrets").
		Set("is_deleted", true).
		Where(sq.Eq{
			"is_deleted": false,
			"id":         ids,
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, q, args...)

	return err
}
