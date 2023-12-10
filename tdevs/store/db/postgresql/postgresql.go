package postgresql

import (
	"tdevs/store"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(drivername string, dSN string) (store.IDriver, error) {
	db, err := sqlx.Connect(drivername, dSN)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}
