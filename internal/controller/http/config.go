package http

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"APP_PORT" required:"true"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
