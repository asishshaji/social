package store

import (
	"context"
	"time"
)

type User struct {
	ID             int32
	Username       string
	HashedPassword string
	Company        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s *Store) CreateUser(ctx context.Context, user *User) (*User, error) {
	user, err := s.driver.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CheckUserExists(ctx context.Context, username string) bool {
	user, _ := s.driver.GetUser(ctx, username)

	return user.Username == username
}

func (s *Store) GetUser(ctx context.Context, username string) (*User, error) {
	return s.driver.GetUser(ctx, username)
}

func (s *Store) AddUserToCache(username string) {
	s.users.Store(username, true)
}
