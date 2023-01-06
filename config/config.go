package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configfile []byte

type Config struct {
	Port     string   `yaml:"port"`
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func New() (*Config, error) {
	var c Config

	err := yaml.Unmarshal(configfile, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil

}
