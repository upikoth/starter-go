package models

type Session struct {
	ID     string
	Token  string
	UserID string
}

type SessionWithUserRole struct {
	Session
	UserRole UserRole
}
