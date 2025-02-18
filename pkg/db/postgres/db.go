package postgres

import (
	"fmt"
	"log"

	"github.com/Oxygenta-Team/FortiKey/pkg/db"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sqlx.DB
}

func CreateStorage(dbConfig *db.Config) (*Storage, error) {
	dsn := dbConfig.DNS()

	connect, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %s", err)
	}

	// TODO: IN THE FUTURE, USE A LOGGER!!!
	log.Println("Successfully connected to the database!")
	return &Storage{DB: connect}, nil
}

func (p *Storage) Close() error {
	return p.DB.Close()
}
