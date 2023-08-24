package emailsender

import (
	"github.com/kelseyhightower/envconfig"
	gomail "gopkg.in/mail.v2"
)

type EmailSender struct {
	from     string
	password string
	host     string
	port     int
}

func New() (*EmailSender, error) {
	config, configErr := newConfig()

	if configErr != nil {
		return nil, configErr
	}

	return &EmailSender{
		from:     config.EmailFrom,
		password: config.EmailPassword,
		host:     config.EmailHost,
		port:     config.EmailPort,
	}, nil
}

func (s *EmailSender) Send(to string, subject string, body string) error {
	d := gomail.NewDialer(s.host, s.port, s.from, s.password)

	m := gomail.NewMessage()

	m.SetHeader("From", s.from)
	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	return d.DialAndSend(m)
}

type config struct {
	EmailFrom     string `envconfig:"EMAIL_FROM" required:"true"`
	EmailPassword string `envconfig:"EMAIL_PASSWORD" required:"true"`
	EmailHost     string `envconfig:"EMAIL_HOST" required:"true"`
	EmailPort     int    `envconfig:"EMAIL_PORT" required:"true"`
}

func newConfig() (*config, error) {
	config := &config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
