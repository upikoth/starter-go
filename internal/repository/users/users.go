package users

import (
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
