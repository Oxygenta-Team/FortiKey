package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

func UnmarshalYAML(path string, config any) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(fileData, &config)
}
