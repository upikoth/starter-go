package oauth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type Oauth struct {
	logger logger.Logger
}

func New(
	log logger.Logger,
) *Oauth {
	return &Oauth{
		logger: log,
	}
}

func (o *Oauth) sendHTTPRequest(
	ctx context.Context,
	method string,
	url string,
	req any,
) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	if ctx.Err() != nil {
		return []byte{}, errors.WithStack(ctx.Err())
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	request, err := http.NewRequestWithContext(
		ctxWithTimeout,
		method,
		url,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(request) //nolint:bodyclose // сделано чуть ниже.

	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	if res.StatusCode != http.StatusOK {
		statusErr := errors.Errorf(
			"не удалось выполнить запрос %s: %s, статус ответа - %d",
			method,
			url,
			res.StatusCode,
		)

		o.logger.Error(statusErr.Error())
		o.logger.Error(string(bodyBytes))

		return []byte{}, statusErr
	}

	return bodyBytes, nil
}
