package v1

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type HandlerV1 struct {
	logger  logger.Logger
	service *service.Service
	config  *config
}

type config struct {
	SiteURL string `envconfig:"SITE_URL" required:"true"`
}

func newConfig() (*config, error) {
	config := &config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

func New(logger logger.Logger, service *service.Service) *HandlerV1 {
	config, configErr := newConfig()

	if configErr != nil {
		log.Fatal(configErr)
	}

	logger.Debug(config.SiteURL)

	return &HandlerV1{
		logger,
		service,
		config,
	}
}
