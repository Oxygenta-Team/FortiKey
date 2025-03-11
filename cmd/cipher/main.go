package cipher

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Oxygenta-Team/FortiKey/pkg/cfg"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
)

// TODO: PATH FROM FLAG WITH DEFAULT PATH FROM ENV?
var defaultPath = "./config.yaml"

// TODO: GET PORT FROM CONFIG
var temporalPort = ":1221"

func main() {
	var config config
	err := cfg.UnmarshalYAML(defaultPath, &config)
	if err != nil {
		logrus.Fatalf("error during creation config, err: %s", err)
		return
	}

	log := logging.NewLogger(config.LogLevel)

	storage, err := pg.CreateStorage(&config.DB)
	if err != nil {
		log.Fatalf("error during creation storage(db), err: %s", err)
		return
	}
	log.Println("Successfully connected to db")

	svc := services.NewServices(postgres.NewRepoManager(), storage, log)

	r := router.NewRouter(svc)

	log.Printf("Server is starting on %s", temporalPort)
	log.Fatal(http.ListenAndServe(temporalPort, r))
}
