package passwordrecoveryrequests

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"gorm.io/gorm"
)

type PasswordRecoveryRequests struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *PasswordRecoveryRequests {
	return &PasswordRecoveryRequests{
		db:     db,
		logger: logger,
	}
}

func (p *PasswordRecoveryRequests) WithTx(tx *gorm.DB) *PasswordRecoveryRequests {
	return &PasswordRecoveryRequests{
		db:     tx,
		logger: p.logger,
	}
}
