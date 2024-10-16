package passwordrecoveryrequests

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

type PasswordRecoveryRequests struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *PasswordRecoveryRequests {
	return &PasswordRecoveryRequests{
		db:     db,
		logger: logger,
	}
}

func (p *PasswordRecoveryRequests) Create(
	inputCtx context.Context,
	passwordRecoveryRequestToCreate models.PasswordRecoveryRequest,
) (models.PasswordRecoveryRequest, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToCreate)
	res := p.db.WithContext(ctx).Create(&passwordRecoveryRequest)
	createdPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return createdPasswordRecoveryRequest, res.Error
	}

	return createdPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) Update(
	inputCtx context.Context,
	passwordRecoveryRequestToUpdate models.PasswordRecoveryRequest,
) (models.PasswordRecoveryRequest, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Update")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToUpdate)
	res := p.db.WithContext(ctx).Updates(&passwordRecoveryRequest)
	updatedPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return updatedPasswordRecoveryRequest, res.Error
	}

	return updatedPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) GetByEmail(
	inputCtx context.Context,
	email string,
) (models.PasswordRecoveryRequest, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.GetByEmail")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbmodels.PasswordRecoveryRequest{}
	res := p.db.
		WithContext(ctx).
		Where(ydbmodels.PasswordRecoveryRequest{Email: email}).
		FirstOrInit(&passwordRecoveryRequest)

	if res.RowsAffected == 0 {
		passwordRecoveryRequest = ydbmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundPasswordRecoveryRequest, res.Error
	}

	return foundPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (models.PasswordRecoveryRequest, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.GetByToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbmodels.PasswordRecoveryRequest{}
	res := p.db.
		WithContext(ctx).
		Where(ydbmodels.PasswordRecoveryRequest{ConfirmationToken: confirmationToken}).
		FirstOrInit(&passwordRecoveryRequest)

	if res.RowsAffected == 0 {
		passwordRecoveryRequest = ydbmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundPasswordRecoveryRequest, res.Error
	}

	return foundPasswordRecoveryRequest, nil
}
