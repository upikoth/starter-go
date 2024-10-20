package ycp

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"net/mail"
)

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
