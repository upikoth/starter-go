package registrations

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"gorm.io/gorm"
)

type Registration struct {
	ID                string `gorm:"primarykey"`
	Email             string
	ConfirmationToken string
}

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

	registration := toLocalModel(registrationToCreate)
	res := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := fromLocalModel(registration)

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

	registration := toLocalModel(registrationToUpdate)
	res := r.db.WithContext(ctx).Updates(&registration)
	updatedRegistration := fromLocalModel(registration)

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

	registration := Registration{}
	res := r.db.WithContext(ctx).Where(Registration{Email: email}).FirstOrInit(&registration)
	foundRegistration := fromLocalModel(registration)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return foundRegistration, res.Error
	}

	return foundRegistration, nil
}

func toLocalModel(registration models.Registration) Registration {
	return Registration{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}

func fromLocalModel(registration Registration) models.Registration {
	return models.Registration{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}
