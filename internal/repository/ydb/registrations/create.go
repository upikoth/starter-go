package registrations

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (r *Registrations) Create(
	inputCtx context.Context,
	registrationToCreate *models.Registration,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Create")
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

	registration := ydbmodels.NewYDBRegistrationModel(*registrationToCreate)
	dbRes := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := registration.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &createdRegistration, nil
}