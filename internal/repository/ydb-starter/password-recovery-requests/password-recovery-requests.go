package passwordrecoveryrequests

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb-starter/ydbs-models"
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
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.PasswordRecoveryRequests.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbsmodels.NewYdbsPasswordRecoveryRequestModel(passwordRecoveryRequestToCreate)
	res := p.db.WithContext(ctx).Create(&passwordRecoveryRequest)
	createdPasswordRecoveryRequest := passwordRecoveryRequest.FromYdbsModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.PasswordRecoveryRequests.Update")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbsmodels.NewYdbsPasswordRecoveryRequestModel(passwordRecoveryRequestToUpdate)
	res := p.db.WithContext(ctx).Updates(&passwordRecoveryRequest)
	updatedPasswordRecoveryRequest := passwordRecoveryRequest.FromYdbsModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.PasswordRecoveryRequests.GetByEmail")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbsmodels.PasswordRecoveryRequest{}
	res := p.db.
		WithContext(ctx).
		Where(ydbsmodels.PasswordRecoveryRequest{Email: email}).
		FirstOrInit(&passwordRecoveryRequest)

	if res.RowsAffected == 0 {
		passwordRecoveryRequest = ydbsmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYdbsModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.PasswordRecoveryRequests.GetByToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := ydbsmodels.PasswordRecoveryRequest{}
	res := p.db.
		WithContext(ctx).
		Where(ydbsmodels.PasswordRecoveryRequest{ConfirmationToken: confirmationToken}).
		FirstOrInit(&passwordRecoveryRequest)

	if res.RowsAffected == 0 {
		passwordRecoveryRequest = ydbsmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundPasswordRecoveryRequest, res.Error
	}

	return foundPasswordRecoveryRequest, nil
}
