package config

import (
	"os"

	"github.com/sherifabdlnaby/configuro"
)

type Config struct {
	Database *Database `validate:"required"`
	Server   struct {
		HTTP *ServerHTTP `validate:"required"`
	}
}

type Database struct {
	Cluster  []string `validate:"required"`
	Keyspace string   `validate:"required"`
	Username string   `validate:"required"`
	Password string   `validate:"required"`
}

type ServerHTTP struct {
	Host     string `validate:"required"`
	Port     int
	UseHTTPS bool
	CertPath string
}

func NewConfig(configPath string) (*Config, error) {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, err
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
		configuro.WithLoadFromEnvVars("APP"),
	)

	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = loader.Load(config)
	if err != nil {
		return nil, err
	}

	err = loader.Validate(config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
