package postgresql

import (
	"fmt"
	"os"
	"tdevs/store"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewDB(drivername string, dSN string) (store.IDriver, error) {
	db, err := sqlx.Connect(drivername, dSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
		return nil, err
	}
	return &PostgresDB{
		db: db,
	}, nil
}
