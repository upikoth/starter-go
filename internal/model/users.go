package model

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusArchived UserStatus = "archived"
)

type User struct {
	tableName    struct{}   `pg:"users"` //nolint:unused // Имя таблицы.
	ID           int        `json:"id" pg:"id"`
	Name         string     `json:"name" pg:"name"`
	Email        string     `json:"email" pg:"email"`
	Status       UserStatus `json:"status" pg:"status"`
	PasswordHash string     `json:"-" pg:"password_hash"`
}
