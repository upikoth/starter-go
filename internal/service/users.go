package service

import "github.com/upikoth/starter-go/internal/model"

func (s *Service) GetUsers() ([]model.User, error) {
	return s.repository.Users.GetAll()
}

func (s *Service) GetUser(id int) (model.User, error) {
	return s.repository.Users.Get(id)
}
