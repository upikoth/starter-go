package sessions

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
)

func (s *Sessions) Create(
	inputCtx context.Context,
	sessionToCreate models.Session,
) (res *models.Session, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.Create")
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

	session := ydbmodels.NewYDBSessionModel(sessionToCreate)
	dbRes := s.db.WithContext(ctx).Create(&session)
	createdSession := session.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &createdSession, nil
}
