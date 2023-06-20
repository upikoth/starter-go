package controller

import (
	"log"

	"github.com/upikoth/starter-go/internal/controller/http"
)

type Controller struct {
	http *http.HTTP
}

func New() *Controller {
	config, configErr := http.NewConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	return &Controller{
		http: http.New(config),
	}
}

func (c *Controller) Start() error {
	return c.http.Start()
}
