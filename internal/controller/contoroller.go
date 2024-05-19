package controller

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type Controller struct {
	http   *http.HTTP
	logger logger.Logger
	config *config.Controller
}

func New(config *config.Config, logger logger.Logger) *Controller {
	return &Controller{
		http:   http.New(&config.Controller.HTTP, logger),
		logger: logger,
		config: &config.Controller,
	}
}

func (c *Controller) Start() error {
	return c.http.Start()
}
