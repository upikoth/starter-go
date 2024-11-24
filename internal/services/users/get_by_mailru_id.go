package users

import (
	"context"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

func (u *Users) GetByMailRuID(
	inputCtx context.Context,
	mailRuID string,
) (*models.User, error) {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
	defer span.End()

	user, err := u.repositories.users.GetByMailRuID(ctx, mailRuID)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return nil, constants.ErrUserNotFound
	}

	if err != nil {
		tracing.HandleError(span, err)
		return nil, err
	}

	return user, nil
}
