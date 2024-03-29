package repository

import (
	"github.com/upikoth/starter-go/internal/repository/pg"
	registrationsRepository "github.com/upikoth/starter-go/internal/repository/registrations"
	sessionsRepository "github.com/upikoth/starter-go/internal/repository/sessions"
	usersRepository "github.com/upikoth/starter-go/internal/repository/users"
)

type Repository struct {
	Users         *usersRepository.Users
	Registrations *registrationsRepository.Registrations
	Sessions      *sessionsRepository.Sessions
	pg            *pg.Pg
}

func New() *Repository {
	pg := pg.New()

	return &Repository{
		Users:         usersRepository.New(pg),
		Registrations: registrationsRepository.New(pg),
		Sessions:      sessionsRepository.New(pg),
		pg:            pg,
	}
}

func (r *Repository) Start() error {
	return r.pg.Connect()
}

func (r *Repository) Stop() error {
	return r.pg.Disconnect()
}
