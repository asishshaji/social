package store

import (
	"context"
	"time"
)

type status string

const (
	ENABLED_STATUS     status = "enabled"
	DISABLED_STATUS    status = "disabled"
	BLACKLISTED_STATUS status = "blacklisted"
)

type User struct {
	ID             int32     `db:"id"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	Status         status    `db:"status"`
	Company        string    `db:"company"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (s *Store) CreateUser(ctx context.Context, user *User) (*User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = ENABLED_STATUS

	err := s.driver.CreateUser(ctx, user)
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
