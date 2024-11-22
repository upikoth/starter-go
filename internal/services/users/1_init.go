package users

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
)

type Users struct {
	repository *repositories.Repository
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	repository *repositories.Repository,
) *Users {
	return &Users{
		repository: repository,
		logger:     logger,
	}
}
