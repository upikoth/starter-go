package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type Registration struct {
	ID                string
	Email             string
	ConfirmationToken string
}

func NewYDBRegistrationModel(registration *models.Registration) *Registration {
	return &Registration{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}

func (r *Registration) FromYDBModel() *models.Registration {
	return &models.Registration{
		ID:                r.ID,
		Email:             r.Email,
		ConfirmationToken: r.ConfirmationToken,
	}
}
