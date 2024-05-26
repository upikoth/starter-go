package service

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service/registrations"
)

type Service struct {
	Registrations *registrations.Registrations
}

func New(
	logger logger.Logger,
	config *config.Config,
	repository *repository.Repository,
) (*Service, error) {
	return &Service{
		Registrations: registrations.New(
			logger,
			&config.Service.Registrations,
			repository,
		),
	}, nil
}
