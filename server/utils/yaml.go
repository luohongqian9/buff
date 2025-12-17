package utils

import (
	"io/fs"
	"os"
	"server/global"

	"gopkg.in/yaml.v3"
)

const configFlie = "./config.yaml"

func LoadYaml() ([]byte, error) {
	return os.ReadFile(configFlie)
}

func SaveYaml() error {
	data, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFlie, data, fs.ModePerm)
}
