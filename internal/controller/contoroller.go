package controller

import (
	"github.com/upikoth/starter-go/internal/controller/http"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type Controller struct {
	http   *http.HTTP
	logger logger.Logger
}

func New(logger logger.Logger) *Controller {
	config, configErr := http.NewConfig()
	if configErr != nil {
		logger.Fatal(configErr)
	}

	return &Controller{
		http:   http.New(config, logger),
		logger: logger,
	}
}

func (c *Controller) Start() error {
	return c.http.Start()
}
