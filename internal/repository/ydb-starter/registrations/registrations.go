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
	registrationInput models.Registration,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Repository: YdbStarter.Registrations.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := toLocalModel(registrationInput)
	res := r.db.WithContext(ctx).Create(&registration)
	createdRegistration := fromLocalModel(registration)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return createdRegistration, res.Error
	}

	return createdRegistration, nil
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
