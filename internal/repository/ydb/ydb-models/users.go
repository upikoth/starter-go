package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type User struct {
	ID           string `gorm:"primarykey"`
	Email        string
	PasswordHash string
	Role         string
}

func NewYDBUserModel(user models.User) User {
	return User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
	}
}

func (u *User) FromYDBModel() models.User {
	return models.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         models.UserRole(u.Role),
	}
}
