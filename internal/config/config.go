package config

import (
	"os"

	"github.com/xh3sh/go-htmx-example/internal/model"
	"gopkg.in/yaml.v2"
)

func LoadConfig() (*model.Config, error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var config model.Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
