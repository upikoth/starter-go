package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	VkID         string
}

func NewYDBUserModel(user *models.User) *User {
	return &User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
		VkID:         user.VkID,
	}
}

func (u *User) FromYDBModel() *models.User {
	return &models.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         models.UserRole(u.Role),
		VkID:         u.VkID,
	}
}
