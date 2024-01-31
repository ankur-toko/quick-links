package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type FileConfig struct{}

const (
	default_config_file = "./properties/config.yml"
)

var defaultFileConfig *FileConfig

func GetConfigProvider() *FileConfig {
	if defaultFileConfig == nil {
		defaultFileConfig = CreateFileConfig()
	}
	return defaultFileConfig
}

func CreateFileConfig() *FileConfig {
	f := FileConfig{}
	v := os.Getenv("QL_CONFIG_FILEPATH")
	if v == "" {
		v = default_config_file
	}
	f.ReloadFile(v)
	return &f
}
func (f *FileConfig) Reload() error {
	v := os.Getenv("QL_CONFIG_FILEPATH")
	if v == "" {
		v = default_config_file
	}
	f.ReloadFile(v)
	return nil
}

func (f *FileConfig) ReloadFile(filename string) error {
	v := filename
	var err error
	v, err = filepath.Abs(v)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, e := os.Stat(v)
	if os.IsNotExist(e) {
		return nil
	}
	config.AddDriver(yaml.Driver)
	err = config.LoadFiles(v)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (fc *FileConfig) Get(key string) (string, bool) {
	data := config.String(key)
	if data != "" {
		return data, true
	} else {
		return "", false
	}
}

func (fc *FileConfig) Set(key string, val string) {
	config.Set(key, val)
}
