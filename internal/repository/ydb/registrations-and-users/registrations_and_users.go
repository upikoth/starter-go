package registrationsandusers

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

type RegistrationsAndUsers struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *RegistrationsAndUsers {
	return &RegistrationsAndUsers{
		db:     db,
		logger: logger,
	}
}

func (r *RegistrationsAndUsers) DeleteRegistrationAndCreateUser(
	inputCtx context.Context,
	registrationToDelete models.Registration,
	userToCreate models.User,
) (models.User, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.RegistrationsAndUsers.DeleteRegistrationAndCreateUser")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.NewYDBRegistrationModel(registrationToDelete)
	user := ydbmodels.NewYDBUserModel(userToCreate)

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&registration).Error; err != nil {
			return err
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	})

	createdUser := user.FromYDBModel()

	if err != nil {
		sentry.CaptureException(err)
		return createdUser, err
	}

	return createdUser, nil
}
