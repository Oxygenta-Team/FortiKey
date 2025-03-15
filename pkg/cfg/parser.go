package cfg

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func UnmarshalYAML(path string, config any) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	logrus.Debug("unmarshal was success")
	return yaml.Unmarshal(fileData, config)
}
