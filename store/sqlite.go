package store

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const createTableQuery string = `
CREATE TABLE IF NOT EXISTS keys (
	id INTEGER NOT NULL PRIMARY KEY,
	key VARCHAR(36) NOT NULL
)`

type SQLite struct {
	FilePath string
	Database *sql.DB
}

func (s *SQLite) Init() error {
	db, err := sql.Open("sqlite", s.FilePath)
	if err != nil {
		return err
	}

	if _, err = db.Exec(createTableQuery); err != nil {
		return err
	}

	s.Database = db

	return nil
}

func (s *SQLite) SetAPIKey(discord_id, api_key string) error {
	_, err := s.Database.Exec("INSERT OR REPLACE INTO keys (id, key) VALUES(?,?)", discord_id, api_key)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLite) GetAPIKey(discord_id string) (string, error) {
	res, err := s.Database.Query("SELECT key FROM keys WHERE id=?", discord_id);
	if err != nil {
		return "", err
	}

	var key string;

	if res.Next() {
		err = res.Scan(&key);
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("couldn't find the key for this user")
	}

	return key, nil;
}

func (s *SQLite) GetIDByKey(api_key string) (string, error) {
	res, err := s.Database.Query("SELECT id FROM keys WHERE key=?", api_key);
	if err != nil {
		return "", err
	}
	var id string;

	if res.Next() {
		err = res.Scan(&id);
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("couldn't find the id for this user")
	}

	return id, nil;
}