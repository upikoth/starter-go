package handler

import (
	"context"
	"errors"
	"net/http"

	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type Handler struct {
	logger  logger.Logger
	service *service.Service
}

func New(logger logger.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) NewError(_ context.Context, err error) *starter.ErrorResponseStatusCode {
	modelErr := &models.Error{}

	code := models.ErrorCodeValidationByOpenapi
	statusCode := http.StatusBadRequest

	if errors.As(err, &modelErr) {
		code = modelErr.Code
		statusCode = modelErr.GetStatusCode()
	}

	return &starter.ErrorResponseStatusCode{
		StatusCode: statusCode,
		Response: starter.ErrorResponse{
			Success: starter.ErrorResponseSuccessFalse,
			Error: starter.ErrorResponseError{
				Code:        string(code),
				Description: err.Error(),
			},
		},
	}
}
