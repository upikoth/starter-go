package sessions

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
)

type Sessions struct {
	repository *repositories.Repository
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	repository *repositories.Repository,
) *Sessions {
	return &Sessions{
		repository: repository,
		logger:     logger,
	}
}
