package registrations

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydbs-models"
	"gorm.io/gorm"
)

type Registrations struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(
	db *gorm.DB,
	logger logger.Logger,
) *Registrations {
	return &Registrations{
		db:     db,
		logger: logger,
	}
}

func (r *Registrations) Create(
	inputCtx context.Context,
	registrationToCreate models.Registration,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Registrations.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbsmodels.NewYdbsRegistrationModel(registrationToCreate)
	res := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := registration.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return createdRegistration, res.Error
	}

	return createdRegistration, nil
}

func (r *Registrations) Update(
	inputCtx context.Context,
	registrationToUpdate models.Registration,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Registrations.Update")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbsmodels.NewYdbsRegistrationModel(registrationToUpdate)
	res := r.db.WithContext(ctx).Updates(&registration)
	updatedRegistration := registration.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return updatedRegistration, res.Error
	}

	return updatedRegistration, nil
}

func (r *Registrations) GetByEmail(
	inputCtx context.Context,
	email string,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Registrations.GetByEmail")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbsmodels.Registration{}
	res := r.db.
		WithContext(ctx).
		Where(ydbsmodels.Registration{Email: email}).
		FirstOrInit(&registration)

	if res.RowsAffected == 0 {
		registration = ydbsmodels.Registration{}
	}
	foundRegistration := registration.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundRegistration, res.Error
	}

	return foundRegistration, nil
}

func (r *Registrations) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Registrations.GetByToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbsmodels.Registration{}
	res := r.db.
		WithContext(ctx).
		Where(ydbsmodels.Registration{ConfirmationToken: confirmationToken}).
		FirstOrInit(&registration)

	if res.RowsAffected == 0 {
		registration = ydbsmodels.Registration{}
	}
	foundRegistration := registration.FromYdbsModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundRegistration, res.Error
	}

	return foundRegistration, nil
}
