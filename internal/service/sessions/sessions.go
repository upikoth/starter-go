package sessions

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Sessions struct {
	repository *repository.Repository
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	repository *repository.Repository,
) *Sessions {
	return &Sessions{
		repository: repository,
		logger:     logger,
	}
}
