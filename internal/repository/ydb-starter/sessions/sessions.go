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

func (r *Sessions) Create(
	inputCtx context.Context,
	sessionToCreate models.Session,
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Sessions.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session := ydbsmodels.NewYdbsSessionModel(sessionToCreate)
	res := r.db.WithContext(ctx).Create(&session)
	createdSession := session.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return createdSession, res.Error
	}

	return createdSession, nil
}
