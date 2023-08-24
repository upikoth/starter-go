package pg

import "github.com/kelseyhightower/envconfig"

type config struct {
	DatabaseName     string `envconfig:"DATABASE_NAME" required:"true"`
	DatabaseAddr     string `envconfig:"DATABASE_ADDR" required:"true"`
	DatabaseUser     string `envconfig:"DATABASE_USER" required:"true"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" required:"true"`
}

func newConfig() (*config, error) {
	config := &config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
