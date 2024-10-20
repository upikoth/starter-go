package sessions

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"gorm.io/gorm"
)

type Sessions struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *Sessions {
	return &Sessions{
		db:     db,
		logger: logger,
	}
}
