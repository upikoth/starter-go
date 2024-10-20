//nolint:dupl // тут нужно дублировать
package passwordrecoveryrequests

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

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
