package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Controller Controller
}

type Controller struct {
	HTTP ControllerHTTP
}

type ControllerHTTP struct {
	Port string `envconfig:"PORT" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
