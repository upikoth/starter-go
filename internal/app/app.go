package app

import (
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type App struct {
	config     *Config
	logger     logger.Logger
	controller *controller.Controller
	repository *repository.Repository
}

func New(config *Config, logger logger.Logger) *App {
	controller := controller.New(logger)
	repository := repository.New()

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
