package main

import (
	"flag"
	"github.com/Oxygenta-Team/FortiKey/pkg/cfg"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
)

// TODO: PATH FROM FLAG WITH DEFAULT PATH FROM ENV?
var defaultConfigPath = "./cmd/cipher/config.yaml"

func main() {
	var configPath string
	flag.StringVar(
		&configPath,
		"config-path",
		defaultConfigPath,
		"provides a path to configuration file with extension .yaml")

	flag.Parse()
	logrus.Println("Config path is:", configPath)
	var config config
	err := cfg.UnmarshalYAML(configPath, &config)
	if err != nil {
		logrus.Fatalf("error during creation config, err: %s", err)
		return
	}

	logrus.Printf("%+v", config)
	level, err := logging.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to parse logging level. Err: %v", err)
	}
	logger, err := logging.NewLogger(level)
	if err != nil {
		logger.Fatalf("error during creation logger, err: %s", err)
		return
	}
	storage, err := pg.CreateStorage(&config.DB)
	if err != nil {
		logger.Fatalf("error during creation storage(db), err: %s", err)
		return
	}
	logger.Println("Successfully connected to db")

	producer := kafka.NewProducer(&config.Kafka)
	svc := services.NewServices(postgres.NewRepoManager(), producer, storage, logger)

	r := router.NewRouter(svc)

	svc.StartConsumer(logger, &config.Kafka)

	logger.Infof("Server is starting on %s", config.Addr)
	//logger.Fatal(http.ListenAndServe(config.Addr, nil))
	logger.Fatal(http.ListenAndServe(config.Addr, r))
}
