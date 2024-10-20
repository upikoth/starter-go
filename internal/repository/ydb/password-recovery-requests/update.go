package passwordrecoveryrequests

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

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
