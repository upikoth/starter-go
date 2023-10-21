package registrations

import (
	"github.com/upikoth/starter-go/internal/constants"
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

func (r *Registrations) GetByToken(token string) (model.Registration, error) {
	registration := model.Registration{
		RegistrationConfirmationToken: token,
	}

	count, err := r.pg.Db.
		Model(&registration).
		Where("registration_confirmation_token = ?", registration.RegistrationConfirmationToken).
		SelectAndCount()

	if err != nil {
		return registration, err
	}

	if count == 0 {
		return registration, constants.ErrDbNotFound
	}

	return registration, nil
}

func (r *Registrations) DeleteByID(id int) error {
	registration := model.Registration{
		ID: id,
	}

	_, err := r.pg.Db.
		Model(&registration).
		WherePK().
		Delete()

	if err != nil {
		return err
	}

	return nil
}
