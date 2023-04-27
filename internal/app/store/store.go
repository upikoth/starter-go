package store

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Store struct {
	db     *pg.DB
	config *Config
}

func New() *Store {
	config := NewConfig()

	return &Store{
		config: config,
	}
}

func (s *Store) Connect() error {
	db := pg.Connect(&pg.Options{
		User:     s.config.DatabaseUser,
		Password: s.config.DatabasePassword,
		Database: s.config.DatabaseName,
		Addr:     s.config.DatabaseAddr,
	})

	err := db.Ping(context.Background())
	if err != nil {
		return nil
	}

	s.db = db

	return nil
}

func (s *Store) Disconnect() error {
	return s.db.Close()
}
