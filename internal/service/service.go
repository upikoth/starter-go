package service

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service/registrations"
)

type Service struct {
	Registrations *registrations.Registrations
	logger        logger.Logger
}

func New(logger logger.Logger) (*Service, error) {
	repository, err := repository.New(logger)

	if err != nil {
		return nil, err
	}

	return &Service{
		Registrations: registrations.New(logger, repository),
		logger:        logger,
	}, nil
}
