package handler

import (
	v1 "github.com/upikoth/starter-go/internal/app/handler/v1"
	"github.com/upikoth/starter-go/internal/app/store"
)

type Handler struct {
	V1 *v1.HandlerV1
}

func New(store *store.Store, jwtSecret []byte) *Handler {
	return &Handler{
		V1: v1.New(store, jwtSecret),
	}
}
