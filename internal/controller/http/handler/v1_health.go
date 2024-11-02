package handler

import (
	"context"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"go.opentelemetry.io/otel"
)

func (h *Handler) V1CheckHealth(inputCtx context.Context) (*app.SuccessResponse, error) {
	tracer := otel.Tracer("Controller: V1CheckHealth")
	_, span := tracer.Start(inputCtx, "Controller: V1CheckHealth")
	defer span.End()

	return &app.SuccessResponse{
		Success: app.SuccessResponseSuccessTrue,
	}, nil
}
