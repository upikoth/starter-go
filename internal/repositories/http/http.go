package http

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/http/oauth"
)

type HTTP struct {
	Oauth *oauth.Oauth
}

func New(
	log logger.Logger,
) (*HTTP, error) {
	return &HTTP{
		Oauth: oauth.New(log),
	}, nil
}
