package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
)

func (h *Handler) V1CheckHealth(ctx context.Context) (*starter.SuccessResponse, error) {
	span := sentry.StartSpan(ctx, "Controller: V1CheckHealth")
	defer func() {
		span.Finish()
	}()

	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}
