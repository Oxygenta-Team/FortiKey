package integration_test

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	ta "github.com/Oxygenta-Team/FortiKey/pkg/testassets"
)

var (
	db *pg.Storage
	ts *httptest.Server
)

const serviceName = "cipher"

func TestMain(m *testing.M) {
	dockerDB, err := ta.CreateDockerDB(serviceName)
	if err != nil {
		log.Fatal(err)
	}
	db = dockerDB

	svc := services.NewServices(postgres.NewRepoManager(), db, ta.Logger)
	r := router.NewRouter(svc)
	ts = httptest.NewServer(r)

	err = initializeTestData()
	if err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func initializeTestData() error {
	return nil // If need some data from another services, init it here
}
