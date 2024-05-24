package registrations

import (
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
)

func (r *Registrations) Create(params models.RegistrationCreateParams) (models.Registration, error) {
	registration := models.Registration{
		ID:                uuid.New().String(),
		Email:             params.Email,
		ConfirmationToken: uuid.New().String(),
	}

	err := r.repository.YcpStarter.SendEmail()

	if err != nil {
		return registration, err
	}

	return registration, nil
}
