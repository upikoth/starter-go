package users

import (
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
)

type Option func(user *models.User)

func newUser(
	options ...Option,
) *models.User {
	user := &models.User{
		ID:   models.UserID(uuid.New().String()),
		Role: models.UserRoleUser,
	}

	for _, option := range options {
		option(user)
	}

	return user
}

func withPasswordHash(ph string) Option {
	return func(user *models.User) {
		user.PasswordHash = ph
	}
}

func withEmail(email string) Option {
	return func(user *models.User) {
		user.Email = email
	}
}
