package users

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (u *Users) Update(
	inputCtx context.Context,
	userToUpdate models.User,
) (res *models.User, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.Update")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		} else {
			bytes, _ := json.Marshal(res)
			span.SetData("Result", string(bytes))
		}
		span.Finish()
	}()
	ctx := span.Context()

	user := ydbmodels.NewYDBUserModel(userToUpdate)
	dbRes := u.db.WithContext(ctx).Updates(&user)
	updatedUser := user.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &updatedUser, nil
}
