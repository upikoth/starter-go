package repository

import (
	"github.com/upikoth/starter-go/internal/repository/pg"
)

type Repository struct {
	pg *pg.Pg
}

func New() *Repository {
	return &Repository{
		pg: pg.New(),
	}
}

func (r *Repository) Start() error {
	return r.pg.Connect()
}

func (r *Repository) Stop() error {
	return r.pg.Disconnect()
}
