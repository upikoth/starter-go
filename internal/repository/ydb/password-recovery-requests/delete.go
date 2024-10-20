package passwordrecoveryrequests

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (p *PasswordRecoveryRequests) Delete(
	inputCtx context.Context,
	id string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Delete")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	dbRes := p.db.
		WithContext(ctx).
		Delete(ydbmodels.PasswordRecoveryRequest{ID: id})

	if dbRes.Error != nil {
		return errors.WithStack(dbRes.Error)
	}

	return nil
}
