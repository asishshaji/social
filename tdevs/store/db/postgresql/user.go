package postgresql

import (
	"context"

	"tdevs/store"
)

func (db *DB) CreateUser(context.Context, *store.User) (*store.User, error) {

	return nil, nil
}

func (db *DB) GetUser(context.Context, string) (*store.User, error) {
	return nil, nil
}
