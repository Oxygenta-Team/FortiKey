package cipher

import (
	"github.com/sirupsen/logrus"

	"github.com/Oxygenta-Team/FortiKey/pkg/db"
)

type config struct {
	Addr     string       `yaml:"addr"`
	LogLevel logrus.Level `yaml:"logLevel"`
	DB       db.Config    `yaml:"db"`
}
