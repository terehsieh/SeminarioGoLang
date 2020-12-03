package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // adding sqlite driver support
	"github.com/terehsieh/SeminarioGoLang/internal/config"
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}
