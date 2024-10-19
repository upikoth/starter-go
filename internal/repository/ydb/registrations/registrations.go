package registrations

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
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
	registrationToCreate *models.Registration,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Create")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.NewYDBRegistrationModel(*registrationToCreate)
	dbRes := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := registration.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &createdRegistration, nil
}

func (r *Registrations) Update(
	inputCtx context.Context,
	registrationToUpdate models.Registration,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.Update")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.NewYDBRegistrationModel(registrationToUpdate)
	dbRes := r.db.WithContext(ctx).Updates(&registration)
	updatedRegistration := registration.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return &updatedRegistration, nil
}

func (r *Registrations) GetByEmail(
	inputCtx context.Context,
	email string,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.GetByEmail")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.Registration{}
	dbRes := r.db.
		WithContext(ctx).
		Where(ydbmodels.Registration{Email: email}).
		FirstOrInit(&registration)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		registration = ydbmodels.Registration{}
	}
	foundRegistration := registration.FromYDBModel()

	return &foundRegistration, nil
}

func (r *Registrations) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (res *models.Registration, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.Registrations.GetByToken")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	registration := ydbmodels.Registration{}
	dbRes := r.db.
		WithContext(ctx).
		Where(ydbmodels.Registration{ConfirmationToken: confirmationToken}).
		FirstOrInit(&registration)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if dbRes.RowsAffected == 0 {
		registration = ydbmodels.Registration{}
	}
	foundRegistration := registration.FromYDBModel()

	return &foundRegistration, nil
}
