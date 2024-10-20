//nolint:dupl // тут нужно дублировать
package sessions

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (s *Sessions) GetByToken(
	inputCtx context.Context,
	token string,
) (res *models.Session, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.GetByToken")
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

	session := ydbmodels.Session{}
	dbRes := s.db.
		WithContext(ctx).
		Where(ydbmodels.Session{Token: token}).
		FirstOrInit(&session)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		session = ydbmodels.Session{}
	}
	foundSession := session.FromYDBModel()

	return &foundSession, nil
}
