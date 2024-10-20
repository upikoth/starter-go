//nolint:dupl // тут нужно дублировать
package registrations

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (r *Registrations) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.GetByToken")
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

	registration := ydbmodels.Registration{}
	dbRes := r.db.
		WithContext(ctx).
		Where(ydbmodels.Registration{ConfirmationToken: confirmationToken}).
		FirstOrInit(&registration)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		registration = ydbmodels.Registration{}
	}
	foundRegistration := registration.FromYDBModel()

	return &foundRegistration, nil
}
