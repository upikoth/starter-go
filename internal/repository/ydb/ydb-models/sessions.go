package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type Session struct {
	ID       string `gorm:"primarykey"`
	UserID   string
	UserRole string
	Token    string
}

func NewYDBSessionModel(session models.Session) Session {
	return Session{
		ID:       session.ID,
		UserID:   session.UserID,
		UserRole: string(session.UserRole),
		Token:    session.Token,
	}
}

func (u *Session) FromYDBModel() models.Session {
	return models.Session{
		ID:       u.ID,
		UserID:   u.UserID,
		UserRole: models.UserRole(u.UserRole),
		Token:    u.Token,
	}
}
