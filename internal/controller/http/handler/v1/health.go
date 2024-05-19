package v1

import (
	"context"

	starterApi "github.com/upikoth/starter-go/internal/generated/starter"
)

func (h *HandlerV1) CheckHealth(_ context.Context) (*starterApi.DefaultSuccessResponse, error) {
	return &starterApi.DefaultSuccessResponse{
		Success: starterApi.DefaultSuccessResponseSuccessTrue,
	}, nil
}
