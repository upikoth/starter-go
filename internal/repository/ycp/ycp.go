package ycp

import (
	"context"
	"net/mail"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	smtpclient "github.com/upikoth/starter-go/internal/pkg/smtp-client"
)

type Ycp struct {
	logger logger.Logger
	config *config.Ycp
	client *smtpclient.SMTPClient
}

func New(
	logger logger.Logger,
	config *config.Ycp,
) (*Ycp, error) {
	client, err := smtpclient.New(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Ycp{
		logger: logger,
		config: config,
		client: client,
	}, nil
}

func (y *Ycp) SendEmail(
	inputCtx context.Context,
	toEmail string,
	title string,
	body string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YCP.SendEmail")
	defer func() {
		if err != nil {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()

	err = y.client.Connect()

	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		disconnectError := y.client.Disconnect()

		if err != nil {
			y.logger.Error(disconnectError.Error())
		}
	}()

	from := mail.Address{
		Name:    y.config.FromName,
		Address: y.config.FromAddress,
	}
	to := mail.Address{
		Name:    "",
		Address: toEmail,
	}

	message := y.client.CreateMessage(from, to, title, body)
	err = y.client.SendEmail(from, to, message)

	return errors.WithStack(err)
}
