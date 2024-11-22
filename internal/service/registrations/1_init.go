package registrations

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Registrations struct {
	logger     logger.Logger
	config     *config.Registrations
	repository *repository.Repository
}

func New(
	logger logger.Logger,
	config *config.Registrations,
	repository *repository.Repository,
) *Registrations {
	return &Registrations{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}
