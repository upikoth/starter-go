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
	return &Service{
		Registrations: registrationsService.New(logger, repository),
		Users:         usersService.New(logger, repository),
		logger:        logger,
		repository:    repository,
	}
}
