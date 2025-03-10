package cipher

import "github.com/Oxygenta-Team/FortiKey/pkg/db"

type config struct {
	Addr string    `yaml:"addr"`
	DB   db.Config `yaml:"db"`
}
