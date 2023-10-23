package model

type Session struct {
	tableName struct{} `pg:"sessions"` //nolint:unused // Имя таблицы.
	ID        int      `json:"id" pg:"id"`
	Token     string   `json:"token" pg:"token"`
	UserID    int      `json:"userId" pg:"user_id"`
	UserAgent string   `json:"userAgent" pg:"user_agent"`
	CreatedAt string   `json:"createdAt" pg:"created_at"`
	ExpiredAt string   `json:"expiredAt" pg:"expired_at"`
}
