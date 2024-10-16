package passwordrecoveryrequestsandusers

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

type PasswordRecoveryRequestsAndUsers struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *PasswordRecoveryRequestsAndUsers {
	return &PasswordRecoveryRequestsAndUsers{
		db:     db,
		logger: logger,
	}
}

func (r *PasswordRecoveryRequestsAndUsers) DeletePasswordRecoveryRequestAndUpdateUser(
	inputCtx context.Context,
	passwordRecoveryRequestToDelete models.PasswordRecoveryRequest,
	userToUpdate models.User,
) (models.User, error) {
	span := sentry.StartSpan(
		inputCtx,
		"Repository: YDB.PasswordRecoveryRequestsAndUsers.DeletePasswordRecoveryRequestAndUpdateUser",
	)
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToDelete)
	user := ydbmodels.NewYDBUserModel(userToUpdate)

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&passwordRecoveryRequest).Error; err != nil {
			return err
		}

		if err := tx.Updates(&user).Error; err != nil {
			return err
		}

		return nil
	})

	updatedUser := user.FromYDBModel()

	if err != nil {
		sentry.CaptureException(err)
		return updatedUser, err
	}

	return updatedUser, nil
}
