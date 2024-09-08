package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	app "github.com/upikoth/starter-go/internal/generated/app"
)

func (h *Handler) V1CheckHealth(ctx context.Context) (*app.SuccessResponse, error) {
	span := sentry.StartSpan(ctx, "Controller: V1CheckHealth")
	defer func() {
		span.Finish()
	}()

	return &app.SuccessResponse{
		Success: app.SuccessResponseSuccessTrue,
	}, nil
}
