package handler

import (
	"context"

	starter "github.com/upikoth/starter-go/internal/generated/starter"
)

func (h *Handler) V1CreateRegistration(
	_ context.Context,
	req *starter.V1RegistrationsCreateRegistrationRequestBody,
) (*starter.SuccessResponse, error) {
	h.logger.Info(req)
	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}
