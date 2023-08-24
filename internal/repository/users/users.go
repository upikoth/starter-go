package users

import (
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/repository/pg"
)

type Users struct {
	pg *pg.Pg
}

func New(pg *pg.Pg) *Users {
	return &Users{
		pg: pg,
	}
}

func (u *Users) GetAll() ([]model.User, error) {
	users := []model.User{}

	err := u.pg.Db.
		Model(&users).
		Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Users) GetByID(id int) (model.User, error) {
	user := model.User{
		ID: id,
	}

	count, err := u.pg.Db.
		Model(&user).
		WherePK().
		SelectAndCount()

	if err != nil {
		return user, err
	}

	if count == 0 {
		return user, constants.ErrDbNotFound
	}

	return user, nil
}

func (u *Users) GetByEmail(email string) (model.User, error) {
	user := model.User{
		Email: email,
	}

	count, err := u.pg.Db.
		Model(&user).
		Where("email = ?", user.Email).
		SelectAndCount()

	if err != nil {
		return user, err
	}

	if count == 0 {
		return user, constants.ErrDbNotFound
	}

	return user, nil
}
