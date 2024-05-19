package v1

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type HandlerV1 struct {
	logger logger.Logger
}

func New(logger logger.Logger) *HandlerV1 {
	return &HandlerV1{
		logger,
	}
}
