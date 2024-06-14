package models

type PasswordRecoveryRequest struct {
	ID                string
	Email             string
	ConfirmationToken string
}
