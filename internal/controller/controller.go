package controller

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/services"
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	http   *http.HTTP
	logger logger.Logger
	config *config.Controller
}

func New(
	config *config.Config,
	logger logger.Logger,
	service *services.Service,
	tp trace.TracerProvider,
) (*Controller, error) {
	httpInstance, err := http.New(&config.Controller.HTTP, logger, service, tp)

	if err != nil {
		return nil, err
	}

	return &Controller{
		http:   httpInstance,
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
