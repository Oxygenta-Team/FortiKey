package models

type User struct {
	ID        int       `db:"id"`
	Login     string    `db:"login"`
	Email     string    `db:"email"`
	IsDeleted bool      `db:"is_deleted"`
}
