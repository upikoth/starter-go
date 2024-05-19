package handler

import (
	"context"

	v1 "github.com/upikoth/starter-go/internal/controller/http/handler/v1"
	starterApi "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type Handler struct {
	v1 *v1.HandlerV1
}

func New(logger logger.Logger) *Handler {
	return &Handler{
		v1: v1.New(logger),
	}
}

func (h *Handler) V1GetHealth(context context.Context) (*starterApi.DefaultSuccessResponse, error) {
	return h.v1.CheckHealth(context)
}

func (h *Handler) NewError(_ context.Context, _ error) *starterApi.DefaultErrorResponseStatusCode {
	return &starterApi.DefaultErrorResponseStatusCode{}
}
