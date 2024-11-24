package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
)

type getMailRuUserInfoResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (o *Oauth) GetMailRuUserInfo(
	ctx context.Context,
	accessToken string,
) (*models.OauthUserInfo, error) {
	bodyBytes, err := o.sendHTTPRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://oauth.mail.ru/userinfo?access_token=%s", accessToken),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getMailRuUserInfoResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.OauthUserInfo{
		ID:    resParsed.ID,
		Email: resParsed.Email,
	}, nil
}
