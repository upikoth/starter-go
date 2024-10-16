package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1CheckCurrentSession(
	inputCtx context.Context,
	params app.V1CheckCurrentSessionParams,
) (*app.SuccessResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1GetCurrentSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

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
	span := sentry.StartSpan(inputCtx, "Controller: V1CreateSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

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
	span := sentry.StartSpan(inputCtx, "Controller: V1DeleteSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	err := h.service.Sessions.DeleteByID(ctx, params.ID)

	if err != nil {
		return nil, err
	}

	return &app.SuccessResponse{
		Success: app.SuccessResponseSuccessTrue,
	}, nil
}
