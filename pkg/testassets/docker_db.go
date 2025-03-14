package testassets

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func buildDBURL(resource *dockertest.Resource, serviceName string) string {
	return fmt.Sprintf("postgres://dev:12345@%v/%s?sslmode=disable", resource.GetHostPort("5432/tcp"), serviceName)
}

func getPostgresContainerConfig(serviceName string) *dockertest.RunOptions {
	return &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.4",
		Env: []string{
			"POSTGRES_USER=dev",
			"POSTGRES_PASSWORD=12345",
			fmt.Sprintf("POSTGRES_DB=%s", serviceName),
			"listen_addresses = '*'",
		},
	}
}

func startPostgresContainer(pool *dockertest.Pool, serviceName string) (*dockertest.Resource, error) {
	runOptions := getPostgresContainerConfig(serviceName)

	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return nil, fmt.Errorf("could not start PostgreSQL container. err: %w", err)
	}

	return resource, nil
}

func getMigrationsDir() (string, error) {
	_, err := os.Stat("../migrations")
	if err == nil {
		return "../migrations", nil
	}

	if os.IsNotExist(err) {
		return "../repository/postgres/migrations", nil
	}

	return "", fmt.Errorf("could not check migrations directory. err: %w", err)
}

func applyMigrations(migrations, dbURL string) error {
	migration, err := migrate.New("file://"+migrations, dbURL)
	if err != nil {
		return fmt.Errorf("could not open migration sources. err: %w", err)
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not apply migrations. err: %w", err)
	}

	return nil
}

func CreateDockerDB(name string) (*postgres.Storage, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("error during creating pool, err: %s", err)
	}

	rs, err := startPostgresContainer(pool, name)
	if err != nil {
		return nil, fmt.Errorf("error during creating container, err: %s, name: %s", err, name)
	}
	dbUrl := buildDBURL(rs, name)

	var db *postgres.Storage
	if err := pool.Retry(func() error {
		db, err = postgres.CreateStorageByURL(dbUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return nil, fmt.Errorf("error during connecting to container, err: %s, name: %s", err, name)
	}

	migrationsDir, err := getMigrationsDir()
	if err != nil {
		return nil, err
	}

	if err = applyMigrations(migrationsDir, dbUrl); err != nil {
		return nil, fmt.Errorf("could not apply migrations. err: %w", err)
	}

	return db, nil
}
