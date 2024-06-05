package ydbsmodels

import "github.com/upikoth/starter-go/internal/models"

type Session struct {
	ID     string `gorm:"primarykey"`
	UserID string
	Token  string
}

func NewYdbsSessionModel(session models.Session) Session {
	return Session{
		ID:     session.ID,
		UserID: session.UserID,
		Token:  session.Token,
	}
}

func (u *Session) FromYdbsModel() models.Session {
	return models.Session{
		ID:     u.ID,
		UserID: u.UserID,
		Token:  u.Token,
	}
}
