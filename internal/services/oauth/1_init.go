package oauth

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

type Oauth struct {
	logger     logger.Logger
	vkConfig   *oauth2.Config
	repository *repositories.Repository
}

func New(
	log logger.Logger,
	cfg config.Oauth,
	repo *repositories.Repository,
) *Oauth {
	return &Oauth{
		logger: log,
		vkConfig: &oauth2.Config{
			ClientID:     cfg.VkClientID,
			ClientSecret: cfg.VkClientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  cfg.VkRedirectURL,
			Scopes:       []string{"email"},
		},
		repository: repo,
	}
}
