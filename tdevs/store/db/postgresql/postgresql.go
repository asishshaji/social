package postgresql

import (
	"tdevs/store"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDB() (store.IDriver, error) {
	return nil, nil
}
