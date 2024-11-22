package app

import (
	"context"
	"fmt"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
	"github.com/upikoth/starter-go/internal/services"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	config     *config.Config
	logger     logger.Logger
	repository *repositories.Repository
	service    *services.Service
	controller *controller.Controller
}

func New(
	config *config.Config,
	logger logger.Logger,
	tp trace.TracerProvider,
) (*App, error) {
	repositoryInstance, err := repositories.New(logger, &config.Repository)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации repository: %s", err))
		return nil, err
	}

	serviceInstance, err := services.New(logger, config, repositoryInstance)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации service: %s", err))
		return nil, err
	}

	controllerInstance, err := controller.New(config, logger, serviceInstance, tp)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации controller: %s", err))
		return nil, err
	}

	return &App{
		config:     config,
		logger:     logger,
		repository: repositoryInstance,
		service:    serviceInstance,
		controller: controllerInstance,
	}, nil
}

func (s *App) Start(ctx context.Context) error {
	err := s.repository.Connect(ctx)

	if err != nil {
		return err
	}

	s.logger.Debug("Подключение к repository прошло без ошибок")

	return s.controller.Start()
}

func (s *App) Stop(ctx context.Context) error {
	err := s.repository.Disconnect(ctx)

	if err != nil {
		return err
	}

	s.logger.Debug("Отключение от repository прошло без ошибок")

	return s.controller.Stop(ctx)
}
