package handler

import (
	"context"

	starter "github.com/upikoth/starter-go/internal/generated/starter"
)

func (h *Handler) V1CheckHealth(_ context.Context) (*starter.SuccessResponse, error) {
	return &starter.SuccessResponse{
		Success: starter.SuccessResponseSuccessTrue,
	}, nil
}
