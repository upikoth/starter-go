package service

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service/registrations"
)

type Service struct {
	Registrations *registrations.Registrations
	logger        logger.Logger
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*Service, error) {
	repository, err := repository.New(logger, &config.Repository)

	if err != nil {
		return nil, err
	}

	return &Service{
		Registrations: registrations.New(
			logger,
			&config.Service.Registrations,
			repository,
		),
		logger: logger,
	}, nil
}
