package oauth

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
	"net/http"
)

func (o *Oauth) GetAuthorizeURL(
	inputCtx context.Context,
	oauthSource models.OauthSource,
) (string, error) {
	tracer := otel.Tracer("Service: Oauth.GetAuthorizeURL")
	_, span := tracer.Start(inputCtx, "Oauth: Sessions.GetAuthorizeURL")
	defer span.End()

	switch oauthSource {
	case models.OauthSourceVk:
		return o.vkConfig.AuthCodeURL(""), nil
	case models.OauthSourceOK:
		return "", nil
	case models.OauthSourceMail:
		return "", nil
	case models.OauthSourceYandex:
		return "", nil
	default:
		return "", &models.Error{
			Code:        models.ErrorCodeOauthSourceNotExist,
			Description: "Incorrect email or password",
			StatusCode:  http.StatusBadRequest,
		}
	}
}
