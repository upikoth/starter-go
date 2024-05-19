package app

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type App struct {
	config     *config.Config
	logger     logger.Logger
	controller *controller.Controller
}

func New(config *config.Config, logger logger.Logger) *App {
	controller := controller.New(config, logger)

	return &App{
		config:     config,
		logger:     logger,
		controller: controller,
	}
}

func (s *App) Start() error {
	return s.controller.Start()
}
