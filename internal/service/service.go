package service

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Service struct {
	logger     logger.Logger
	repository *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Service {
	return &Service{
		logger,
		repository,
	}
}
