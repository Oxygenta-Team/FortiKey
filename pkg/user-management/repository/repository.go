package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"

	sq "github.com/Masterminds/squirrel"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(login, email string) (*models.User, error) {
	query := `INSERT INTO users(login, email) VALUES ($1, $2) RETURNING id, created_at`
	user := &models.User{Login: login, Email: email}
	err := r.db.QueryRowx(query, login, email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func (r *Repository) DeleteUser(id int) error {
	query := `UPDATE "user" SET deleted=true WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) RestoreUser(id int) error {
	query := `UPDATE "user" SET deleted=false WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, login, email, created_at, is_delete FROM "user" WHERE id = $1 AND is_delete = FALSE`
	user := &models.User{}
	err := r.db.Get(user, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}
	return user, nil
}
