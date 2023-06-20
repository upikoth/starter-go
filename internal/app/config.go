package app

import "github.com/kelseyhightower/envconfig"

type Config struct {
	JwtSecret []byte `envconfig:"APP_JWT_SECRET" required:"true"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
