package controller

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type Controller struct {
	http   *http.HTTP
	logger logger.Logger
	config *config.Controller
}

func New(config *config.Config, logger logger.Logger) (*Controller, error) {
	http, err := http.New(&config.Controller.HTTP, logger)

	if err != nil {
		return nil, err
	}

	return &Controller{
		http:   http,
		logger: logger,
		config: &config.Controller,
	}, nil
}

func (c *Controller) Start() error {
	return c.http.Start()
}

func (c *Controller) Stop(ctx context.Context) error {
	return c.http.Stop(ctx)
}
