package users

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb-starter/ydbs-models"
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
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Users.GetByEmail")
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
