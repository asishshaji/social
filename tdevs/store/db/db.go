package db

import (
	"errors"
	p "tdevs/server/profile"
	"tdevs/store"
	"tdevs/store/db/postgresql"
)

func NewDBDriver(profile *p.Profile) (store.IDriver, error) {
	var driver store.IDriver
	var err error

	switch profile.Driver {
	case p.POSTGRESQL:
		{
			driver, err = postgresql.NewDB()
		}
	}

	if err != nil {
		return nil, errors.New("invalid driver")
	}

	return driver, nil
}
