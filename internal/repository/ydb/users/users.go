package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

type Users struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *Users {
	return &Users{
		db:     db,
		logger: logger,
	}
}

func (u *Users) GetByEmail(
	inputCtx context.Context,
	email string,
) (res *models.User, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.GetByEmail")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
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

func (u *Users) Update(
	inputCtx context.Context,
	userToUpdate models.User,
) (res *models.User, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.Update")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
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

func (u *Users) GetList(
	inputCtx context.Context,
	params models.UsersGetListParams,
) (res *models.UserList, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Users.GetList")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
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
