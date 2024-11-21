package handler

import (
	"context"
	"net/http"

	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
)

func (h *Handler) V1AuthorizeUsingOauth(
	inputCtx context.Context,
	req *app.V1AuthorizeUsingOauthRequestBody,
) (*app.V1AuthorizeUsingOauthResponse, error) {
	tracer := otel.Tracer("Controller: V1AuthorizeUsingOauth")
	ctx, span := tracer.Start(inputCtx, "Controller: V1AuthorizeUsingOauth")
	defer span.End()

	var oauthSource models.OauthSource

	switch req.OauthSource {
	case app.V1AuthorizeUsingOauthRequestBodyOauthSourceVk:
		oauthSource = models.OauthSourceVk
	case app.V1AuthorizeUsingOauthRequestBodyOauthSourceMail:
		oauthSource = models.OauthSourceMail
	case app.V1AuthorizeUsingOauthRequestBodyOauthSourceOk:
		oauthSource = models.OauthSourceOK
	case app.V1AuthorizeUsingOauthRequestBodyOauthSourceYandex:
		oauthSource = models.OauthSourceYandex
	default:
		return nil, &models.Error{
			Code:        models.ErrorCodeOauthSourceNotExist,
			Description: "Incorrect email or password",
			StatusCode:  http.StatusBadRequest,
		}
	}

	url, err := h.service.Oauth.GetAuthorizeURL(ctx, oauthSource)

	if err != nil {
		return nil, err
	}

	return &app.V1AuthorizeUsingOauthResponse{
		Success: true,
		Data: app.V1AuthorizeUsingOauthResponseData{
			URL: url,
		},
	}, nil
}
