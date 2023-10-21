package users

import (
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Users struct {
	logger     logger.Logger
	repository *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Users {
	return &Users{
		logger,
		repository,
	}
}

func (u *Users) GetAll() ([]model.User, error) {
	return u.repository.Users.GetAll()
}

func (u *Users) Get(id int) (model.User, error) {
	return u.repository.Users.GetByID(id)
}

func (u *Users) GetByEmail(email string) (model.User, error) {
	return u.repository.Users.GetByEmail(email)
}

func (u *Users) Create(user model.User) error {
	return u.repository.Users.Create(user)
}
