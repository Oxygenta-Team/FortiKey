package main

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/db"
	"github.com/Oxygenta-Team/FortiKey/pkg/queue/kafka"
)

type config struct {
	Addr     string       `yaml:"addr"`
	LogLevel string       `yaml:"logLevel"`
	DB       db.Config    `yaml:"database"`
	Kafka    kafka.Config `yaml:"kafka"`
}
