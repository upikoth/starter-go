package passwordrecoveryrequests

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type PasswordRecoveryRequests struct {
	repository *repository.Repository
	logger     logger.Logger
	config     *config.PasswordRecoveryRequests
}

func New(
	logger logger.Logger,
	config *config.PasswordRecoveryRequests,
	repository *repository.Repository,
) *PasswordRecoveryRequests {
	return &PasswordRecoveryRequests{
		repository: repository,
		logger:     logger,
		config:     config,
	}
}
