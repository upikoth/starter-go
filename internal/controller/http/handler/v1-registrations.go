package handler

import (
	"context"

	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1CreateRegistration(
	context context.Context,
	req *starter.V1RegistrationsCreateRegistrationRequestBody,
) (*starter.V1RegistrationsCreateRegistrationResponse, error) {
	registrationCreateParams := models.RegistrationCreateParams{
		Email: req.Email,
	}

	registration, err := h.service.Registrations.Create(context, registrationCreateParams)

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
