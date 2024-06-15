//nolint:dupl // тут нужно дублировать
package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1CreatePasswordRecoveryRequest(
	inputCtx context.Context,
	req *starter.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody,
) (*starter.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1CreatePasswordRecoveryRequest")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequestCreateParams := models.PasswordRecoveryRequestCreateParams{
		Email: req.Email,
	}

	passwordRecoveryRequest, err := h.service.PasswordRecoveryRequests.Create(ctx, passwordRecoveryRequestCreateParams)

	if err != nil {
		return nil, err
	}

	return &starter.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse{
		Success: true,
		Data: starter.V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData{
			ID:    passwordRecoveryRequest.ID,
			Email: passwordRecoveryRequest.Email,
		},
	}, nil
}

func (h *Handler) V1ConfirmPasswordRecoveryRequest(
	inputCtx context.Context,
	req *starter.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody,
) (*starter.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1ConfirmPasswordRecoveryRequest")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequestConfirmParams := models.PasswordRecoveryRequestConfirmParams{
		ConfirmationToken: req.ConfirmationToken,
		NewPassword:       string(req.NewPassword),
	}

	session, err := h.service.PasswordRecoveryRequests.Confirm(ctx, passwordRecoveryRequestConfirmParams)

	if err != nil {
		return nil, err
	}

	return &starter.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse{
		Success: true,
		Data: starter.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData{
			Session: starter.V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseDataSession{
				ID:       session.ID,
				Token:    session.Token,
				UserRole: starter.UserRole(session.UserRole),
			},
		},
	}, nil
}
