package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1CheckCurrentSession(
	inputCtx context.Context,
	params starter.V1CheckCurrentSessionParams,
) (*starter.SuccessResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1GetCurrentSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	_, err := h.service.Sessions.CheckToken(ctx, params.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}

func (h *Handler) V1CreateSession(
	inputCtx context.Context,
	req *starter.V1SessionsCreateSessionRequestBody,
) (*starter.V1SessionsCreateSessionResponse, error) {
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

	return &starter.V1SessionsCreateSessionResponse{
		Success: true,
		Data: starter.V1SessionsCreateSessionResponseData{
			Session: starter.Session{
				ID:       session.ID,
				Token:    session.Token,
				UserRole: starter.UserRole(session.UserRole),
			},
		},
	}, nil
}

func (h *Handler) V1DeleteSession(
	inputCtx context.Context,
	params starter.V1DeleteSessionParams,
) (*starter.SuccessResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1DeleteSession")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	err := h.service.Sessions.DeleteByID(ctx, params.ID)

	if err != nil {
		return nil, err
	}

	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}
