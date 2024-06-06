package sessions

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb-starter/ydbs-models"
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
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Sessions.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session := ydbsmodels.NewYdbsSessionModel(sessionToCreate)
	res := s.db.WithContext(ctx).Create(&session)
	createdSession := session.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return createdSession, res.Error
	}

	return createdSession, nil
}

func (s *Sessions) GetByToken(
	inputCtx context.Context,
	token string,
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Sessions.GetByToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session := ydbsmodels.Session{}
	res := s.db.
		WithContext(ctx).
		Where(ydbsmodels.Session{Token: token}).
		FirstOrInit(&session)

	if res.RowsAffected == 0 {
		session = ydbsmodels.Session{}
	}
	foundSession := session.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundSession, res.Error
	}

	return foundSession, nil
}

func (s *Sessions) GetByID(
	inputCtx context.Context,
	id string,
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Sessions.GetByID")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session := ydbsmodels.Session{}
	res := s.db.
		WithContext(ctx).
		Where(ydbsmodels.Session{ID: id}).
		FirstOrInit(&session)

	if res.RowsAffected == 0 {
		session = ydbsmodels.Session{}
	}
	foundSession := session.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundSession, res.Error
	}

	return foundSession, nil
}

func (s *Sessions) DeleteByID(
	inputCtx context.Context,
	id string,
) error {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Sessions.DeleteByID")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	res := s.db.
		WithContext(ctx).
		Delete(ydbsmodels.Session{ID: id})

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return res.Error
	}

	return nil
}
