package oauth

import (
	"context"
	"net/http"

	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

func (o *Oauth) GetAuthorizeURL(
	inputCtx context.Context,
	oauthSource models.OauthSource,
) (string, error) {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	_, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
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
			Code:        models.ErrCodeOauthSourceNotExist,
			Description: "Incorrect email or password",
			StatusCode:  http.StatusBadRequest,
		}
	}
}
