package pg

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseName     string `envconfig:"DATABASE_NAME" required:"true"`
	DatabaseAddr     string `envconfig:"DATABASE_ADDR" required:"true"`
	DatabaseUser     string `envconfig:"DATABASE_USER" required:"true"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" required:"true"`
}

func newConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
