package store

import (
	"context"
	"time"
)

type User struct {
	ID        int32
	Username  string
	Password  string
	Company   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Store) CreateUser(ctx context.Context, user *User) (*User, error) {
	user, err := s.driver.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
