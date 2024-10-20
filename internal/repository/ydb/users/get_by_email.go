package users

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (u *Users) GetByEmail(
	inputCtx context.Context,
	email string,
) (res *models.User, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.GetByEmail")
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

	user := ydbmodels.User{}
	dbRes := u.db.
		WithContext(ctx).
		Where(ydbmodels.User{Email: email}).
		FirstOrInit(&user)
	foundUser := user.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &foundUser, nil
}
