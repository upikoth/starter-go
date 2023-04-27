package store

import (
	"github.com/upikoth/starter-go/internal/app/constants"
	"github.com/upikoth/starter-go/internal/app/model"
)

func (s *Store) GetUsers() ([]model.User, error) {
	users := []model.User{}

	err := s.db.Model(&users).Select()

	if err != nil {
		return users, constants.ErrUsersGetDbError
	}

	return users, nil
}
