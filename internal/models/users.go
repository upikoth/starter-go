package models

type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	UserRole     UserRole
}
