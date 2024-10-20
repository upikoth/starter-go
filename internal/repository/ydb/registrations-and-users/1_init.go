package registrationsandusers

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"gorm.io/gorm"
)

type RegistrationsAndUsers struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *RegistrationsAndUsers {
	return &RegistrationsAndUsers{
		db:     db,
		logger: logger,
	}
}
