package passwordrecoveryrequests

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

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
