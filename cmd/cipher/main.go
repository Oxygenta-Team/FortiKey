package cipher

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"log"
	"net/http"

	"github.com/Oxygenta-Team/FortiKey/pkg/cfg"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
)

// TODO: PATH FROM FLAG WITH DEFAULT PATH FROM ENV?
var defaultPath = "./config.yaml"

// TODO: GET PORT FROM CONFIG
var temporalPort = ":1221"

func main() {
	dbConfig, err := cfg.ParseDBConfig(defaultPath)
	if err != nil {
		log.Fatalf("error during creation config, err: %s", err)
		return
	}

	storage, err := pg.CreateStorage(dbConfig)
	if err != nil {
		log.Fatalf("error during creation storage(db), err: %s", err)
		return
	}
	log.Println("Successfully connected to db")

	svc := services.NewServices(postgres.NewRepoManager(), storage)

	r := router.NewRouter(svc)

	log.Printf("Server is starting on %s", temporalPort)
	log.Fatal(http.ListenAndServe(temporalPort, r))
}
