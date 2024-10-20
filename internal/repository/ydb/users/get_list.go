package users

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (u *Users) GetList(
	inputCtx context.Context,
	params models.UsersGetListParams,
) (res *models.UserList, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.GetList")
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

	var users []ydbmodels.User
	total := int64(0)

	dbRes := u.db.
		WithContext(ctx).
		Model(ydbmodels.User{}).
		Count(&total)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	dbRes = u.db.
		WithContext(ctx).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	var resUsers []models.User
	for _, user := range users {
		resUsers = append(resUsers, user.FromYDBModel())
	}

	return &models.UserList{
		Users: resUsers,
		Total: int(total),
	}, nil
}
