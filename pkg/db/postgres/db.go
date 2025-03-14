package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Oxygenta-Team/FortiKey/pkg/db"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sqlx.DB
}

func CreateStorageByURL(dns string) (*Storage, error) {
	connect, err := sqlx.Connect("postgres", dns)
	if err != nil {
		// TODO return a standart errors
		return nil, fmt.Errorf("failed to connect to the database: %s", err)
	}

	return &Storage{DB: connect}, nil
}

func CreateStorage(dbConfig *db.Config) (*Storage, error) {
	dns := dbConfig.DNS()

	return CreateStorageByURL(dns)
}

func (p *Storage) Close() error {
	return p.DB.Close()
}
