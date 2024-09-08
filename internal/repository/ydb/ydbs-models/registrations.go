package ydbsmodels

import "github.com/upikoth/starter-go/internal/models"

type Registration struct {
	ID                string `gorm:"primarykey"`
	Email             string
	ConfirmationToken string
}

func NewYdbsRegistrationModel(registration models.Registration) Registration {
	return Registration{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}

func (r *Registration) FromYdbsModel() models.Registration {
	return models.Registration{
		ID:                r.ID,
		Email:             r.Email,
		ConfirmationToken: r.ConfirmationToken,
	}
}
