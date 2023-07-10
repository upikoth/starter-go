package pg

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

type Pg struct {
	db     *pg.DB
	config *Config
}

func New() *Pg {
	config, configErr := newConfig()

	if configErr != nil {
		log.Fatal(configErr)
	}

	return &Pg{
		config: config,
	}
}

func (p *Pg) Connect() error {
	db := pg.Connect(&pg.Options{
		User:     p.config.DatabaseUser,
		Password: p.config.DatabasePassword,
		Database: p.config.DatabaseName,
		Addr:     p.config.DatabaseAddr,
	})

	err := db.Ping(context.Background())
	if err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *Pg) Disconnect() error {
	return p.db.Close()
}
