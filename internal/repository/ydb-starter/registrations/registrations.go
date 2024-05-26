package registrations

import (
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

func (r *Registrations) Create(registrationInput models.Registration) (models.Registration, error) {
	registration := toLocalModel(registrationInput)
	res := r.db.Create(&registration)
	createdRegistration := fromLocalModel(registration)

	if res.Error != nil {
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
