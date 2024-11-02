//nolint:dupl // тут нужно дублировать
package handler

import (
	"context"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
)

func (h *Handler) V1CreatePasswordRecoveryRequest(
	inputCtx context.Context,
	req *app.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody,
) (*app.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, error) {
	tracer := otel.Tracer("Controller: V1CreatePasswordRecoveryRequest")
	ctx, span := tracer.Start(inputCtx, "Controller: V1CreatePasswordRecoveryRequest")
	defer span.End()

	passwordRecoveryRequestCreateParams := models.PasswordRecoveryRequestCreateParams{
		Email: req.Email,
	}

	passwordRecoveryRequest, err := h.service.PasswordRecoveryRequests.Create(ctx, passwordRecoveryRequestCreateParams)

	if err != nil {
		return nil, err
	}

	return &app.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse{
		Success: true,
		Data: app.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData{
			ID:    passwordRecoveryRequest.ID,
			Email: passwordRecoveryRequest.Email,
		},
	}, nil
}

func (h *Handler) V1ConfirmPasswordRecoveryRequest(
	inputCtx context.Context,
	req *app.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody,
) (*app.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, error) {
	tracer := otel.Tracer("Controller: V1ConfirmPasswordRecoveryRequest")
	ctx, span := tracer.Start(inputCtx, "Controller: V1ConfirmPasswordRecoveryRequest")
	defer span.End()

	passwordRecoveryRequestConfirmParams := models.PasswordRecoveryRequestConfirmParams{
		ConfirmationToken: req.ConfirmationToken,
		NewPassword:       string(req.NewPassword),
	}

	session, err := h.service.PasswordRecoveryRequests.Confirm(ctx, passwordRecoveryRequestConfirmParams)

	if err != nil {
		return nil, err
	}

	return &app.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse{
		Success: true,
		Data: app.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData{
			Session: app.Session{
				ID:       session.ID,
				Token:    session.Token,
				UserRole: app.UserRole(session.UserRole),
			},
		},
	}, nil
}
