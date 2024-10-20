package sessions

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (s *Sessions) Delete(
	inputCtx context.Context,
	id string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.Delete")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	dbRes := s.db.
		WithContext(ctx).
		Delete(ydbmodels.Session{ID: id})

	if dbRes.Error != nil {
		return errors.WithStack(dbRes.Error)
	}

	return nil
}
