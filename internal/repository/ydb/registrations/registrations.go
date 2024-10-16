package registrations

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
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
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.NewYDBRegistrationModel(registrationToCreate)
	res := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := registration.FromYDBModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Update")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.NewYDBRegistrationModel(registrationToUpdate)
	res := r.db.WithContext(ctx).Updates(&registration)
	updatedRegistration := registration.FromYDBModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.GetByEmail")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.Registration{}
	res := r.db.
		WithContext(ctx).
		Where(ydbmodels.Registration{Email: email}).
		FirstOrInit(&registration)

	if res.RowsAffected == 0 {
		registration = ydbmodels.Registration{}
	}
	foundRegistration := registration.FromYDBModel()

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
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.GetByToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.Registration{}
	res := r.db.
		WithContext(ctx).
		Where(ydbmodels.Registration{ConfirmationToken: confirmationToken}).
		FirstOrInit(&registration)

	if res.RowsAffected == 0 {
		registration = ydbmodels.Registration{}
	}
	foundRegistration := registration.FromYDBModel()

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundRegistration, res.Error
	}

	return foundRegistration, nil
}
