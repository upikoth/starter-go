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

func (s *Store) GetUserById(id int) (model.User, error) {
	user := model.User{
		Id: id,
	}

	count, err := s.db.Model(&user).WherePK().SelectAndCount()

	if count == 0 {
		return model.User{}, constants.ErrUserGetNotFoundById
	}

	if err != nil {
		return model.User{}, constants.ErrUserGetDbError
	}

	return user, nil
}

func (s *Store) CreateUser(user model.User) (model.User, error) {
	result, err := s.db.Model(&user).OnConflict("DO NOTHING").Insert()

	if result.RowsAffected() == 0 {
		return model.User{}, constants.ErrUserPostEmailExist
	}

	if err != nil {
		return model.User{}, constants.ErrUserPostDbError
	}

	return user, nil
}

func (s *Store) DeleteUser(id int) error {
	user := model.User{
		Id: id,
	}

	result, err := s.db.Model(&user).WherePK().Delete()

	if result.RowsAffected() == 0 {
		return constants.ErrUserDeleteNotFoundById
	}

	if err != nil {
		return constants.ErrUserDeleteDbError
	}

	return nil
}

func (s *Store) PatchUser(user model.User) (model.User, error) {
	count, err := s.db.Model(&user).WherePK().Count()

	if err != nil {
		return model.User{}, constants.ErrUserPatchDbError
	}

	if count == 0 {
		return model.User{}, constants.ErrUserPatchNotFoundById
	}

	result, err := s.db.Model(&user).WherePK().OnConflict("DO NOTHING").Update()

	if result == nil {
		return model.User{}, constants.ErrUserPatchEmailExist
	}

	if err != nil {
		return model.User{}, constants.ErrUserPatchDbError
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (model.User, error) {
	user := model.User{
		Email: email,
	}

	count, err := s.db.Model(&user).Where("email = ?", user.Email).SelectAndCount()

	if err != nil {
		return model.User{}, constants.ErrUserGetByEmailDbError
	}

	if count == 0 {
		return model.User{}, constants.ErrUserGetByEmailUserNotExist
	}

	return user, err
}
