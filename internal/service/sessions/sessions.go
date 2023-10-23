package sessions

import (
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
)

type Sessions struct {
	logger     logger.Logger
	repository *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Sessions {
	return &Sessions{
		logger,
		repository,
	}
}

func (s *Sessions) GetAll() ([]model.Session, error) {
	return s.repository.Sessions.GetAll()
}

func (s *Sessions) Create(session model.Session) error {
	return s.repository.Sessions.Create(session)
}
