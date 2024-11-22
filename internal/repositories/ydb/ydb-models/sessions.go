package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type Session struct {
	ID       string
	Token    string
	UserID   string
	UserRole string
}

func NewYDBSessionModel(session *models.Session) *Session {
	return &Session{
		ID:     session.ID,
		Token:  session.Token,
		UserID: session.UserID,
	}
}

func (s *Session) FromYDBModel() *models.SessionWithUserRole {
	return &models.SessionWithUserRole{
		Session: models.Session{
			ID:     s.ID,
			Token:  s.Token,
			UserID: s.UserID,
		},
		UserRole: models.UserRole(s.UserRole),
	}
}
