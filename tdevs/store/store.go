package store

import (
	"tdevs/server/profile"
)

type Store struct {
	driver IDriver
}

func NewStore(driver IDriver, profile *profile.Profile) *Store {
	return &Store{
		driver: driver,
	}
}

func (s *Store) Close() error {
	return nil
}
