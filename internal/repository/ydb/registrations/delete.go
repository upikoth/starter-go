package registrations

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (r *Registrations) Delete(
	inputCtx context.Context,
	id string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Delete")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	dbRes := r.db.
		WithContext(ctx).
		Delete(ydbmodels.Registration{ID: id})

	if dbRes.Error != nil {
		return errors.WithStack(dbRes.Error)
	}

	return nil
}
