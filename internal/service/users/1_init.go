package users

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Users struct {
	repository *repository.Repository
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	repository *repository.Repository,
) *Users {
	return &Users{
		repository: repository,
		logger:     logger,
	}
}
