package http

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/http/oauthmailru"
	"github.com/upikoth/starter-go/internal/repositories/http/oauthyandex"
	"go.opentelemetry.io/otel/trace"
)

type HTTP struct {
	OauthMailRu *oauthmailru.OauthMailRu
	OauthYandex *oauthyandex.OauthYandex
}

func New(
	log logger.Logger,
	cfg *config.HTTP,
	tp trace.TracerProvider,
) (*HTTP, error) {
	oauthMailRu, err := oauthmailru.New(log, &cfg.OauthMailRu, tp)

	if err != nil {
		return nil, err
	}

	oauthYandex, err := oauthyandex.New(log, &cfg.OauthYandex, tp)

	if err != nil {
		return nil, err
	}

	return &HTTP{
		OauthMailRu: oauthMailRu,
		OauthYandex: oauthYandex,
	}, nil
}
