package handler

import (
	"context"
	"errors"
	"net/http"

	app "github.com/upikoth/starter-go/internal/generated/app"
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

func (h *Handler) NewError(_ context.Context, err error) *app.ErrorResponseStatusCode {
	modelErr := &models.Error{}

	code := models.ErrorCodeValidationByOpenapi
	statusCode := http.StatusBadRequest

	if errors.As(err, &modelErr) {
		code = modelErr.Code
		statusCode = modelErr.GetStatusCode()
	}

	return &app.ErrorResponseStatusCode{
		StatusCode: statusCode,
		Response: app.ErrorResponse{
			Success: app.ErrorResponseSuccessFalse,
			Error: app.ErrorResponseError{
				Code:        string(code),
				Description: err.Error(),
			},
		},
	}
}
