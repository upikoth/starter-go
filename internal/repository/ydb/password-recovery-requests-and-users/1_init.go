package passwordrecoveryrequestsandusers

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"gorm.io/gorm"
)

type PasswordRecoveryRequestsAndUsers struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *PasswordRecoveryRequestsAndUsers {
	return &PasswordRecoveryRequestsAndUsers{
		db:     db,
		logger: logger,
	}
}
