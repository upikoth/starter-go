package handler

import (
	"context"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

func (h *Handler) V1CheckCurrentSession(
	inputCtx context.Context,
	params app.V1CheckCurrentSessionParams,
) (*app.SuccessResponse, error) {
	tracer := otel.Tracer(tracing.GetHandlerTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetHandlerTraceName())
	defer span.End()

	_, err := h.service.Sessions.CheckToken(ctx, params.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return &app.SuccessResponse{
		Success: app.SuccessResponseSuccessTrue,
	}, nil
}

func (h *Handler) V1CreateSession(
	inputCtx context.Context,
	req *app.V1SessionsCreateSessionRequestBody,
) (*app.V1SessionsCreateSessionResponse, error) {
	tracer := otel.Tracer(tracing.GetHandlerTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetHandlerTraceName())
	defer span.End()

	sessionCreateParams := models.SessionCreateParams{
		Email:    req.Email,
		Password: string(req.Password),
	}

	session, err := h.service.Sessions.Create(ctx, sessionCreateParams)

	if err != nil {
		return nil, err
	}

	return &app.V1SessionsCreateSessionResponse{
		Success: true,
		Data: app.V1SessionsCreateSessionResponseData{
			Session: app.Session{
				ID:       session.ID,
				Token:    session.Token,
				UserRole: app.UserRole(session.UserRole),
			},
		},
	}, nil
}

func (h *Handler) V1DeleteSession(
	inputCtx context.Context,
	params app.V1DeleteSessionParams,
) (*app.SuccessResponse, error) {
	tracer := otel.Tracer(tracing.GetHandlerTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetHandlerTraceName())
	defer span.End()

	err := h.service.Sessions.DeleteByID(ctx, params.ID)

	if err != nil {
		return nil, err
	}

	return &app.SuccessResponse{
		Success: app.SuccessResponseSuccessTrue,
	}, nil
}
