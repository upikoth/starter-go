package registrationsandusers

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

func (r *RegistrationsAndUsers) DeleteRegistrationAndCreateUser(
	inputCtx context.Context,
	registrationToDelete models.Registration,
	userToCreate models.User,
) (res *models.User, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.RegistrationsAndUsers.DeleteRegistrationAndCreateUser")
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

	registration := ydbmodels.NewYDBRegistrationModel(registrationToDelete)
	user := ydbmodels.NewYDBUserModel(userToCreate)

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if errD := tx.Delete(&registration).Error; errD != nil {
			return errors.WithStack(errD)
		}

		if errC := tx.Create(&user).Error; errC != nil {
			return errors.WithStack(errC)
		}

		return nil
	})

	createdUser := user.FromYDBModel()

	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
