package registrations

import (
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/repository/pg"
)

type Registrations struct {
	pg *pg.Pg
}

func New(pg *pg.Pg) *Registrations {
	return &Registrations{
		pg: pg,
	}
}

func (r *Registrations) Create(registration model.Registration) error {
	_, err := r.pg.Db.
		Model(&registration).
		OnConflict("(email) DO UPDATE").
		Insert()

	if err != nil {
		return err
	}

	return nil
}
