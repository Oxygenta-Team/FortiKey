package cfg

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Oxygenta-Team/FortiKey/pkg/db"
)

// PROBLEM HERE
func ParseDBConfig(path string) (*db.Config, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var settings struct {
		Database db.Config `yaml:"database"`
	}

	if err := yaml.Unmarshal(fileData, &settings); err != nil {
		return nil, err
	}

	return &settings.Database, nil
}
