package handler

import (
	v1 "github.com/upikoth/starter-go/internal/app/handler/v1"
)

type Handler struct {
	V1 *v1.HandlerV1
}

func New() *Handler {
	return &Handler{
		V1: v1.New(),
	}
}
