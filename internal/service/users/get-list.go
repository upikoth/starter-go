package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
)

func (u *Users) GetList(
	inputCtx context.Context,
	params models.UsersGetListParams,
) (models.UserList, error) {
	span := sentry.StartSpan(inputCtx, "Service: Users.GetList")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	return u.repository.Ydb.Users.GetList(ctx, params)
}
