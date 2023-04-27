package v1

import "github.com/upikoth/starter-go/internal/app/store"

type HandlerV1 struct {
	store *store.Store
}

func New(store *store.Store) *HandlerV1 {
	return &HandlerV1{

		store: store,
	}
}
