package handler

import (
	"context"
	"errors"
	"github.com/upikoth/starter-go/internal/config"
	"net/http"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/services"
)

type Handler struct {
	logger  logger.Logger
	service *services.Service
	cfg     *config.ControllerHTTP
}

func New(
	log logger.Logger,
	srv *services.Service,
	cfg *config.ControllerHTTP,
) *Handler {
	return &Handler{
		logger:  log,
		service: srv,
		cfg:     cfg,
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
