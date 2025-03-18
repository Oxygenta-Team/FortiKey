package integration_test

import (
	"github.com/go-redis/redismock/v9"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
	redis_mock "github.com/Oxygenta-Team/FortiKey/pkg/db/redis/mock"
	kafka_mock "github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka/mock"
	ta "github.com/Oxygenta-Team/FortiKey/pkg/testassets"
)

var (
	db    *pg.Storage
	ts    *httptest.Server
	rmock redismock.ClientMock
)

const serviceName = "cipher"

func TestMain(m *testing.M) {
	dockerDB, err := ta.CreateDockerDB(serviceName)
	if err != nil {
		log.Fatal(err)
	}
	db = dockerDB
	rdb, rdbmock := redis_mock.NewMockRedisClient()
	rmock = rdbmock
	svc := services.NewServices(postgres.NewRepoManager(), kafka_mock.NewProducerMock(nil), db, rdb, ta.Logger)
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
