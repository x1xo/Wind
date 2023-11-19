package store

import (
	"fmt"

	"github.com/x1xo/wind/utils"
)

type Store interface {
	Init() error
	SetAPIKey(discord_id, api_key string) error
	GetAPIKey(discord_id string) (string, error)
	GetIDByKey(api_key string) (string, error)
}

var store Store

func GetStore() *Store {
	return &store
}

func Init(config *utils.Config) error {
	if config.Database.Type == "sqlite" {
		s := SQLite{
			FilePath: config.Database.URL,
		}

		err := s.Init()
		if err != nil {
			return err
		}

		store = &s

		return nil
	}

	return fmt.Errorf("database type is not supported")
}
