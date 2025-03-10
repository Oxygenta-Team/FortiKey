package db

import (
	"os"

	"github.com/Oxygenta-Team/FortiKey/pkg/db/postgres"

	"gopkg.in/yaml.v3"
)

func ParseDBConfig(path string) (*postgres.Database, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
		
	var settings struct {
		Database postgres.Database `yaml:"database"`
	}

	if err := yaml.Unmarshal(fileData, &settings); err != nil {
		return nil, err
	}

	return &settings.Database, nil
}
