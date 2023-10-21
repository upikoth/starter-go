package model

type Registration struct {
	tableName                     struct{} `pg:"registrations"` //nolint:unused // Имя таблицы.
	ID                            int      `json:"id" pg:"id"`
	Name                          string   `json:"name" pg:"name"`
	Email                         string   `json:"email" pg:"email"`
	PasswordHash                  string   `json:"-" pg:"password_hash"`
	RegistrationConfirmationToken string   `json:"-" pg:"registration_confirmation_token"`
}
