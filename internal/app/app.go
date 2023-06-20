package app

import "github.com/upikoth/starter-go/internal/controller"

type App struct {
	config     *Config
	controller *controller.Controller
}

func New(config *Config) *App {
	controller := controller.New()

	return &App{
		config:     config,
		controller: controller,
	}
}

func (s *App) Start() error {
	return s.controller.Start()
}
