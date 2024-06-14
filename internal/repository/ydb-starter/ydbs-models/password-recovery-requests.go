package ydbsmodels

import "github.com/upikoth/starter-go/internal/models"

type PasswordRecoveryRequest struct {
	ID                string `gorm:"primarykey"`
	Email             string
	ConfirmationToken string
}

func NewYdbsPasswordRecoveryRequestModel(registration models.PasswordRecoveryRequest) PasswordRecoveryRequest {
	return PasswordRecoveryRequest{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}

func (r *PasswordRecoveryRequest) FromYdbsModel() models.PasswordRecoveryRequest {
	return models.PasswordRecoveryRequest{
		ID:                r.ID,
		Email:             r.Email,
		ConfirmationToken: r.ConfirmationToken,
	}
}
