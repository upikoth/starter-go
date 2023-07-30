package v1

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type HandlerV1 struct {
	logger  logger.Logger
	service *service.Service
}

func New(logger logger.Logger, service *service.Service) *HandlerV1 {
	return &HandlerV1{
		logger,
		service,
	}
}
