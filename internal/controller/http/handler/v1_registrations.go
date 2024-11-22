//nolint:dupl // тут нужно дублировать
package handler

import (
	"context"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

func (h *Handler) V1CreateRegistration(
	inputCtx context.Context,
	req *app.V1RegistrationsCreateRegistrationRequestBody,
) (*app.V1RegistrationsCreateRegistrationResponse, error) {
	tracer := otel.Tracer(tracing.GetHandlerTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetHandlerTraceName())
	defer span.End()

	registrationCreateParams := models.RegistrationCreateParams{
		Email: req.Email,
	}

	registration, err := h.service.Registrations.Create(ctx, registrationCreateParams)

	if err != nil {
		return nil, err
	}

	return &app.V1RegistrationsCreateRegistrationResponse{
		Success: true,
		Data: app.V1RegistrationsCreateRegistrationResponseData{
			ID:    registration.ID,
			Email: registration.Email,
		},
	}, nil
}

func (h *Handler) V1ConfirmRegistration(
	inputCtx context.Context,
	req *app.V1RegistrationsConfirmRegistrationRequestBody,
) (*app.V1RegistrationsConfirmRegistrationResponse, error) {
	tracer := otel.Tracer(tracing.GetHandlerTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetHandlerTraceName())
	defer span.End()

	registrationConfirmParams := models.RegistrationConfirmParams{
		ConfirmationToken: req.ConfirmationToken,
		Password:          string(req.Password),
	}

	session, err := h.service.Registrations.Confirm(ctx, registrationConfirmParams)

	if err != nil {
		return nil, err
	}

	return &app.V1RegistrationsConfirmRegistrationResponse{
		Success: true,
		Data: app.V1RegistrationsConfirmRegistrationResponseData{
			Session: app.Session{
				ID:       session.ID,
				Token:    session.Token,
				UserRole: app.UserRole(session.UserRole),
			},
		},
	}, nil
}
