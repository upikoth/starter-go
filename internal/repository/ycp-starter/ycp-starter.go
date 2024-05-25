package ycpstarter

import (
	"net/mail"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	smtpclient "github.com/upikoth/starter-go/internal/pkg/smtp-client"
)

type YcpStarter struct {
	logger logger.Logger
	config *config.YcpStarter
	client *smtpclient.SMTPClient
}

func New(
	logger logger.Logger,
	config *config.YcpStarter,
) (*YcpStarter, error) {
	client, err := smtpclient.New(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
	)

	if err != nil {
		return nil, err
	}

	return &YcpStarter{
		logger: logger,
		config: config,
		client: client,
	}, nil
}

func (y *YcpStarter) SendEmail(
	toEmail string,
	title string,
	body string,
) error {
	err := y.client.Connect()

	if err != nil {
		return err
	}

	defer func() {
		disconnectError := y.client.Disconnect()

		if err != nil {
			y.logger.Error(disconnectError)
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

	return err
}
