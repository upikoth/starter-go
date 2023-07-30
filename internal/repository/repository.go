package repository

import (
	"github.com/upikoth/starter-go/internal/repository/pg"
	"github.com/upikoth/starter-go/internal/repository/users"
)

type Repository struct {
	Users *users.Users
	pg    *pg.Pg
}

func New() *Repository {
	pg := pg.New()

	return &Repository{
		Users: users.New(pg),
		pg:    pg,
	}
}

func (r *Repository) Start() error {
	return r.pg.Connect()
}

func (r *Repository) Stop() error {
	return r.pg.Disconnect()
}
