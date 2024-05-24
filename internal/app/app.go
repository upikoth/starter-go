package app

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type App struct {
	config     *config.Config
	logger     logger.Logger
	controller *controller.Controller
}

func New(config *config.Config, logger logger.Logger) (*App, error) {
	service, err := service.New(logger)

	if err != nil {
		return nil, err
	}

	controller, err := controller.New(config, logger, service)

	if err != nil {
		return nil, err
	}

	return &App{
		config:     config,
		logger:     logger,
		controller: controller,
	}, nil
}

func (s *App) Start() error {
	return s.controller.Start()
}

func (s *App) Stop(ctx context.Context) error {
	return s.controller.Stop(ctx)
}
