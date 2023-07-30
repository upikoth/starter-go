package service

import "github.com/upikoth/starter-go/internal/model"

func (s *Service) GetUsers() ([]model.User, error) {
	return s.repository.Users.GetAll()
}
