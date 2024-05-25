package registrations

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Registrations struct {
	repository *repository.Repository
	logger     logger.Logger
	config     *config.Registrations
}

func New(
	logger logger.Logger,
	config *config.Registrations,
	repository *repository.Repository,
) *Registrations {
	return &Registrations{
		repository: repository,
		logger:     logger,
		config:     config,
	}
}
