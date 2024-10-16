package app

import (
	"context"
	"fmt"

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
	repositoryInstance, err := repository.New(logger, &config.Repository)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации repository: %s", err))
		return nil, err
	}

	serviceInstance, err := service.New(logger, config, repositoryInstance)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации service: %s", err))
		return nil, err
	}

	controllerInstance, err := controller.New(config, logger, serviceInstance)

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

func (s *App) Start() error {
	err := s.repository.Connect()

	if err != nil {
		s.logger.Error(fmt.Sprintf("Ошибка при подключении к repository: %s", err))
		return err
	}

	s.logger.Debug("Подключение к repository прошло без ошибок")

	return s.controller.Start()
}

func (s *App) Stop(ctx context.Context) error {
	err := s.repository.Disconnect()

	if err != nil {
		s.logger.Error(fmt.Sprintf("Ошибка при отключении от repository: %s", err))
		return err
	}

	s.logger.Debug("Отключение от repository прошло без ошибок")

	return s.controller.Stop(ctx)
}
