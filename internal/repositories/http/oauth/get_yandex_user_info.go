package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
)

type getYandexUserInfoResponse struct {
	ID           string `json:"id"`
	DefaultEmail string `json:"default_email"`
}

func (o *Oauth) GetYandexUserInfo(
	ctx context.Context,
	accessToken string,
) (*models.OauthUserInfo, error) {
	bodyBytes, err := o.sendHTTPRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://login.yandex.ru/info?format=json&oauth_token=%s", accessToken),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getYandexUserInfoResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.OauthUserInfo{
		ID:    resParsed.ID,
		Email: resParsed.DefaultEmail,
	}, nil
}
