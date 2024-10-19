package sessions

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"gorm.io/gorm"
)

type Sessions struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *Sessions {
	return &Sessions{
		db:     db,
		logger: logger,
	}
}

func (s *Sessions) Create(
	inputCtx context.Context,
	sessionToCreate models.Session,
) (res *models.Session, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.Create")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
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

func (s *Sessions) GetByToken(
	inputCtx context.Context,
	token string,
) (res *models.Session, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.GetByToken")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
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

func (s *Sessions) GetByID(
	inputCtx context.Context,
	id string,
) (res *models.Session, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.GetByID")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	session := ydbmodels.Session{}
	dbRes := s.db.
		WithContext(ctx).
		Where(ydbmodels.Session{ID: id}).
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

func (s *Sessions) DeleteByID(
	inputCtx context.Context,
	id string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Sessions.DeleteByID")
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
