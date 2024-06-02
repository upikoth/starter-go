package handler

import (
	"context"

	"github.com/getsentry/sentry-go"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1CreateRegistration(
	inputCtx context.Context,
	req *starter.V1RegistrationsCreateRegistrationRequestBody,
) (*starter.V1RegistrationsCreateRegistrationResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1CreateRegistration")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registrationCreateParams := models.RegistrationCreateParams{
		Email: req.Email,
	}

	registration, err := h.service.Registrations.Create(ctx, registrationCreateParams)

	if err != nil {
		return nil, err
	}

	return &starter.V1RegistrationsCreateRegistrationResponse{
		Success: true,
		Data: starter.V1RegistrationsCreateRegistrationResponseData{
			ID:    registration.ID,
			Email: registration.Email,
		},
	}, nil
}
