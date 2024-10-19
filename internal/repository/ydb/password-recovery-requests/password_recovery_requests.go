package passwordrecoveryrequests

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
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
) (res *models.PasswordRecoveryRequest, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Create")
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

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToCreate)
	dbRes := p.db.WithContext(ctx).Create(&passwordRecoveryRequest)
	createdPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &createdPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) Update(
	inputCtx context.Context,
	passwordRecoveryRequestToUpdate models.PasswordRecoveryRequest,
) (res *models.PasswordRecoveryRequest, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Update")
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

	passwordRecoveryRequest := ydbmodels.NewYDBPasswordRecoveryRequestModel(passwordRecoveryRequestToUpdate)
	dbRes := p.db.WithContext(ctx).Updates(&passwordRecoveryRequest)
	updatedPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &updatedPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) GetByEmail(
	inputCtx context.Context,
	email string,
) (res *models.PasswordRecoveryRequest, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.GetByEmail")
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

	passwordRecoveryRequest := ydbmodels.PasswordRecoveryRequest{}
	dbRes := p.db.
		WithContext(ctx).
		Where(ydbmodels.PasswordRecoveryRequest{Email: email}).
		FirstOrInit(&passwordRecoveryRequest)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		passwordRecoveryRequest = ydbmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	return &foundPasswordRecoveryRequest, nil
}

func (p *PasswordRecoveryRequests) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (res *models.PasswordRecoveryRequest, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.GetByToken")
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

	passwordRecoveryRequest := ydbmodels.PasswordRecoveryRequest{}
	dbRes := p.db.
		WithContext(ctx).
		Where(ydbmodels.PasswordRecoveryRequest{ConfirmationToken: confirmationToken}).
		FirstOrInit(&passwordRecoveryRequest)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		passwordRecoveryRequest = ydbmodels.PasswordRecoveryRequest{}
	}
	foundPasswordRecoveryRequest := passwordRecoveryRequest.FromYDBModel()

	return &foundPasswordRecoveryRequest, nil
}
