package registrations

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Registrations struct {
	repository *repository.Repository
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	repository *repository.Repository,
) *Registrations {
	return &Registrations{
		repository: repository,
		logger:     logger,
	}
}
