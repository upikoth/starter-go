package passwordrecoveryrequestsandusers

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"

	"gorm.io/gorm"
)

func (r *PasswordRecoveryRequestsAndUsers) Delete(
	inputCtx context.Context,
	passwordRecoveryRequestToDelete models.PasswordRecoveryRequest,
	userToUpdate models.User,
) (res *models.User, err error) {
	span := sentry.StartSpan(
		inputCtx,
		"Repository: YDB.PasswordRecoveryRequestsAndUsers.Delete",
	)
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

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToDelete)
	user := ydbmodels.NewYDBUserModel(userToUpdate)

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if errD := tx.Delete(&passwordRecoveryRequest).Error; errD != nil {
			return errors.WithStack(errD)
		}

		if errU := tx.Updates(&user).Error; errU != nil {
			return errors.WithStack(errU)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	updatedUser := user.FromYDBModel()

	return &updatedUser, nil
}
