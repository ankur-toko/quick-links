package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

const (
	default_config_file = "./properties/config.yml"
)

func init() {
	InitializeConfiguration(default_config_file)
}

func InitializeConfiguration(filepath string) error {
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles(filepath)
	return err
}

func Get(key string) (string, bool) {
	data := config.String(key)
	if data != "" {
		return data, true
	} else {
		return "", false
	}
}

func GetValOrDefault(key string) string {
	return config.String(key)
}
