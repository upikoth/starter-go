package model

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	RegisteredAt time.Time `json:"registeredAt"`
	LastVisitAt  time.Time `json:"lastVisitAt"`
}
