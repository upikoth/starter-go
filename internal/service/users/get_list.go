package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
)

func (u *Users) GetList(
	inputCtx context.Context,
	params *models.UsersGetListParams,
) (*models.UserList, error) {
	tracer := otel.Tracer("Service: Users.GetList")
	ctx, span := tracer.Start(inputCtx, "Service: Users.GetList")
	defer span.End()

	res, err := u.repository.YDB.Users.GetList(ctx, params)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeUsersGetListDBError,
			Description: err.Error(),
		}
	}

	return res, nil
}
