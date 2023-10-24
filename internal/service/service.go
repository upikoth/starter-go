package service

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	registrationsService "github.com/upikoth/starter-go/internal/service/registarations"
	sessionsService "github.com/upikoth/starter-go/internal/service/sessions"
	usersService "github.com/upikoth/starter-go/internal/service/users"
)

type Service struct {
	Registrations *registrationsService.Registrations
	Users         *usersService.Users
	Sessions      *sessionsService.Sessions
	logger        logger.Logger
	repository    *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Service {
	usersServiceInstance := usersService.New(logger, repository)
	registrationsServiceInstance := registrationsService.New(logger, repository, usersServiceInstance)
	sessionsServiceInstance := sessionsService.New(logger, repository, usersServiceInstance)

	return &Service{
		Registrations: registrationsServiceInstance,
		Users:         usersServiceInstance,
		Sessions:      sessionsServiceInstance,
		logger:        logger,
		repository:    repository,
	}
}
