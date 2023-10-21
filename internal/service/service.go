package service

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	registrationsService "github.com/upikoth/starter-go/internal/service/registarations"
	usersService "github.com/upikoth/starter-go/internal/service/users"
)

type Service struct {
	Registrations *registrationsService.Registrations
	Users         *usersService.Users
	logger        logger.Logger
	repository    *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Service {
	usersServiceInstance := usersService.New(logger, repository)
	registrationsServiceInstance := registrationsService.New(logger, repository, usersServiceInstance)

	return &Service{
		Registrations: registrationsServiceInstance,
		Users:         usersServiceInstance,
		logger:        logger,
		repository:    repository,
	}
}
