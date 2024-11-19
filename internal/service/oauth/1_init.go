package oauth

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

type Oauth struct {
	logger   logger.Logger
	vkConfig *oauth2.Config
}

func New(
	logger logger.Logger,
	cfg config.Oauth,
) *Oauth {
	return &Oauth{
		logger: logger,
		vkConfig: &oauth2.Config{
			ClientID:     cfg.VkClientID,
			ClientSecret: cfg.VkClientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  cfg.VkRedirectURL,
		},
	}
}
