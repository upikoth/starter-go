package app

import (
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service"
)

type App struct {
	config     *Config
	logger     logger.Logger
	controller *controller.Controller
	repository *repository.Repository
}

func New(config *Config, logger logger.Logger) *App {
	repository := repository.New()
	service := service.New(logger, repository)
	controller := controller.New(logger, service)

	return &App{
		config:     config,
		logger:     logger,
		controller: controller,
		repository: repository,
	}
}

func (s *App) Start() error {
	if repositoryError := s.repository.Start(); repositoryError != nil {
		return repositoryError
	}

	s.logger.Info("Repository успешно подключен")

	defer func() {
		disconnectErr := s.repository.Stop()

		if disconnectErr != nil {
			s.logger.Error(disconnectErr)
		}
	}()

	return s.controller.Start()
}
