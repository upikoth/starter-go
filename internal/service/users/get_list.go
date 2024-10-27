package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
)

func (u *Users) GetList(
	inputCtx context.Context,
	params *models.UsersGetListParams,
) (*models.UserList, error) {
	span := sentry.StartSpan(inputCtx, "Service: Users.GetList")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	res, err := u.repository.YDB.Users.GetList(ctx, params)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeUsersGetListDBError,
			Description: err.Error(),
		}
	}

	return res, nil
}
