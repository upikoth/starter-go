package sessions

import (
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/repository/pg"
)

type Sessions struct {
	pg *pg.Pg
}

func New(pg *pg.Pg) *Sessions {
	return &Sessions{
		pg: pg,
	}
}

func (s *Sessions) GetAll() ([]model.Session, error) {
	sessions := []model.Session{}

	err := s.pg.Db.
		Model(&sessions).
		Select()

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Sessions) Create(session model.Session) error {
	_, err := s.pg.Db.
		Model(&session).
		Insert()

	if err != nil {
		return err
	}

	return nil
}
