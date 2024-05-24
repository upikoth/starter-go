package ycpstarter

import "github.com/upikoth/starter-go/internal/pkg/logger"

type YcpStarter struct {
	logger logger.Logger
}

func New(logger logger.Logger) *YcpStarter {
	return &YcpStarter{
		logger: logger,
	}
}

func (y *YcpStarter) SendEmail() error {
	y.logger.Info("SendEmail was called")
	return nil
}
