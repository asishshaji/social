package util

import (
	"fmt"
	"io"
	"os"

	"github.com/jmoiron/sqlx"
)

type DBOpts struct {
	DB_TYPE     string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	SSLMODE     string
}

func ConnectToDB(opts DBOpts) (*sqlx.DB, error) {
	db, err := sqlx.Connect(opts.DB_TYPE,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", opts.DB_HOST,
			opts.DB_PORT, opts.DB_USER, opts.DB_PASSWORD, opts.DB_NAME, opts.SSLMODE))
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	// db.SetMaxOpenConns()
	// db.SetMaxIdleConns(c.MaxIdle)
	// db.SetConnMaxLifetime(c.MaxLifetime)

	err = installSchema(db)
	if err != nil {
		return nil, fmt.Errorf("error creating schema %s", err)
	}
	return db, nil
}

func installSchema(db *sqlx.DB) error {
	file, err := os.Open("db/schema.sql")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(data)); err != nil {
		return err
	}

	return nil

}
