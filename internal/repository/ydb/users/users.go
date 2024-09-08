package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydbs-models"
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
) (models.User, error) {
	span := sentry.StartSpan(inputCtx, "Repository: Ydb.Users.GetByEmail")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	user := ydbsmodels.User{}
	res := u.db.
		WithContext(ctx).
		Where(ydbsmodels.User{Email: email}).
		FirstOrInit(&user)
	foundUser := user.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundUser, res.Error
	}

	return foundUser, nil
}

func (u *Users) Update(
	inputCtx context.Context,
	userToUpdate models.User,
) (models.User, error) {
	span := sentry.StartSpan(inputCtx, "Repository: Ydb.Users.Update")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	user := ydbsmodels.NewYdbsUserModel(userToUpdate)
	res := u.db.WithContext(ctx).Updates(&user)
	updatedUser := user.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return updatedUser, res.Error
	}

	return updatedUser, nil
}

func (u *Users) GetList(
	inputCtx context.Context,
	params models.UsersGetListParams,
) (models.UserList, error) {
	span := sentry.StartSpan(inputCtx, "Repository: Ydb.Users.GetList")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	users := []ydbsmodels.User{}
	total := int64(0)

	res := u.db.
		WithContext(ctx).
		Model(ydbsmodels.User{}).
		Count(&total)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return models.UserList{}, res.Error
	}

	res = u.db.
		WithContext(ctx).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return models.UserList{}, res.Error
	}

	resUsers := []models.User{}

	for _, user := range users {
		resUsers = append(resUsers, user.FromYdbsModel())
	}

	return models.UserList{
		Users: resUsers,
		Total: int(total),
	}, nil
}
