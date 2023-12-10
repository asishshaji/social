package store

import (
	"sync"
	"tdevs/server/profile"
)

type Store struct {
	driver IDriver
	users  sync.Map
}

func NewStore(driver IDriver, profile *profile.Profile) *Store {
	return &Store{
		driver: driver,
	}
}

func (s *Store) Close() error {
	return nil
}
