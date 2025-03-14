package integration_test

import (
	"context"
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"log"
	"os"
	"testing"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	ta "github.com/Oxygenta-Team/FortiKey/pkg/testassets"
)

var db *pg.Storage

const serviceName = "cipher"

func TestMain(m *testing.M) {
	dockerDB, err := ta.CreateDockerDB(serviceName)
	if err != nil {
		log.Fatal(err)
	}
	db = dockerDB

	exitCode := 0
	defer func() {
		//if err = dockerDB.Pool.Purge(dockerDB.Resource); err != nil {
		//	log.Fatalf("could not purge Postgres resource: %s", err)
		//}
		os.Exit(exitCode)
	}()
	exitCode = m.Run()
}

func TestInsertSecret(t *testing.T) {
	repo := postgres.NewSecretRepo(db)
	ctx := context.Background()

	testCases := []struct {
		name        string
		input       []*models.Secret
		expectError bool
	}{
		{
			name:        "OK",
			input:       []*models.Secret{ta.Secret1, ta.Secret2},
			expectError: false,
		},
		{
			name:        "OK_same",
			input:       []*models.Secret{ta.Secret1},
			expectError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.InsertSecret(ctx, tc.input)
			assert.Equal(t, tc.expectError, err != nil)
		})
	}
}

func TestSelectSecrets(t *testing.T) {
	repo := postgres.NewSecretRepo(db)
	ctx := context.Background()

	testCases := []struct {
		name  string
		input uint64

		expectResponse *models.Secret
		expectError    error
	}{
		{
			name:           "OK",
			input:          ta.Secret1.ID,
			expectResponse: ta.Secret1,
		},
		{
			name:           "OK",
			input:          ta.Secret2.ID,
			expectResponse: ta.Secret2,
		},
		{
			name:        "NotFound",
			input:       ta.Secret3.ID,
			expectError: sql.ErrNoRows,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			secret, err := repo.GetSecretByID(ctx, tc.input)

			if !cmp.Equal(tc.expectResponse, secret, cmpopts.IgnoreFields(models.Secret{}, "Value")) {
				t.Errorf("GetSecretByID() mismatch, expected: %+v, actual: %+v", tc.expectResponse, secret)
			}
			assert.Equal(t, tc.expectError, err)
		})
	}
}

func TestDeleteSecret(t *testing.T) {
	repo := postgres.NewSecretRepo(db)
	ctx := context.Background()

	testCases := []struct {
		name        string
		input       []uint64
		expectError error
	}{
		{
			name:  "OK",
			input: []uint64{ta.Secret1.ID},
		},
		{
			name:  "OK_same",
			input: []uint64{ta.Secret1.ID},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.DeleteSecret(ctx, tc.input)
			assert.Equal(t, tc.expectError, err)
		})
	}
}

func TestSelectSecretsByKey(t *testing.T) {
	repo := postgres.NewSecretRepo(db)
	ctx := context.Background()

	testCases := []struct {
		name  string
		input string

		expectResponse *models.Secret
		expectError    error
	}{
		{
			name:        "OK",
			input:       ta.Secret1.Key,
			expectError: sql.ErrNoRows,
		},
		{
			name:           "NotFound",
			input:          ta.Secret2.Key,
			expectResponse: ta.Secret2,
		},
		{
			name:        "NotFound",
			input:       ta.Secret3.Key,
			expectError: sql.ErrNoRows,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			secret, err := repo.GetSecretByKey(ctx, tc.input)

			if !cmp.Equal(tc.expectResponse, secret, cmpopts.IgnoreFields(models.Secret{}, "Value")) {
				t.Errorf("GetSecretByID() mismatch, \n expected: %+v, \n actual: %+v", tc.expectResponse, secret)
			}
			assert.Equal(t, tc.expectError, err)
		})
	}
}
