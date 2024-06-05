package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
)

func (h *Handler) V1GetCurrentSession(
	inputCtx context.Context,
	params starter.V1GetCurrentSessionParams,
) (*starter.SuccessResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1GetCurrentSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	_, err := h.service.Sessions.CheckSessionToken(ctx, params.AuthorizationToken)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}
