package app

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service"
)

type App struct {
	config     *config.Config
	logger     logger.Logger
	repository *repository.Repository
	service    *service.Service
	controller *controller.Controller
}

func New(config *config.Config, logger logger.Logger) (*App, error) {
	repository, err := repository.New(logger, &config.Repository)

	if err != nil {
		return nil, err
	}

	service, err := service.New(logger, config, repository)

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
		repository: repository,
		service:    service,
		controller: controller,
	}, nil
}

func (s *App) Start() error {
	err := s.repository.Connect()

	if err != nil {
		return err
	}

	return s.controller.Start()
}

func (s *App) Stop(ctx context.Context) error {
	err := s.repository.Disconnect()

	if err != nil {
		return err
	}

	return s.controller.Stop(ctx)
}
