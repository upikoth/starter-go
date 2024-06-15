package ydbsmodels

import "github.com/upikoth/starter-go/internal/models"

type User struct {
	ID           string `gorm:"primarykey"`
	Email        string
	PasswordHash string
	UserRole     string
}

func NewYdbsUserModel(user models.User) User {
	return User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		UserRole:     string(user.UserRole),
	}
}

func (u *User) FromYdbsModel() models.User {
	return models.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		UserRole:     models.UserRole(u.UserRole),
	}
}
