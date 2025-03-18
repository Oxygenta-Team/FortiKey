package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Oxygenta-Team/FortiKey/pkg/cfg"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/repository/postgres"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/router"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"
	"github.com/Oxygenta-Team/FortiKey/pkg/db/redis"
	"github.com/Oxygenta-Team/FortiKey/pkg/logging"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka"

	pg "github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"
)

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
	logger := logging.NewLogger(level)

	storage, err := pg.CreateStorage(&config.DB)
	if err != nil {
		logger.Fatalf("error during creation storage(db), err: %s", err)
		return
	}
	logger.Println("Successfully connected to db")

	redisClient, err := redis.NewClient(&config.Redis)
	if err != nil {
		logger.Fatalf("error during creation redis client, %s", err)
		return
	}

	producer := kafka.NewProducer(&config.Kafka)
	svc := services.NewServices(postgres.NewRepoManager(), producer, storage, redisClient, logger)

	r := router.NewRouter(svc)

	svc.StartConsumer(logger, &config.Kafka)

	logger.Infof("Server is starting on %s", config.Addr)
	logger.Fatal(http.ListenAndServe(config.Addr, r))
}
