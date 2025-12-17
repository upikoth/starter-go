package controller

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	httpcontroller "github.com/upikoth/starter-go/internal/controller/http-controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/services"
	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	http   *httpcontroller.HTTP
	logger logger.Logger
	config *config.Controller
}

func New(
	cfg *config.Config,
	log logger.Logger,
	srvs *services.Services,
	tp trace.TracerProvider,
) (*Controller, error) {
	httpInstance, err := httpcontroller.New(&cfg.Controller.HTTP, log, srvs, tp)
	if err != nil {
		return nil, err
	}

	return &Controller{
		http:   httpInstance,
		logger: log,
		config: &cfg.Controller,
	}, nil
}

func (c *Controller) Start() error {
	return c.http.Start()
}

func (c *Controller) Stop(ctx context.Context) error {
	return c.http.Stop(ctx)
}
