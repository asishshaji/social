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
	ID             int32     `db:"id" json:"id"`
	Username       string    `db:"username" json:"username"`
	Bio            string    `db:"bio" json:"bio"`
	Avatar         string    `db:"avatar" json:"avatar"`
	HashedPassword string    `db:"hashed_password" json:"-"`
	Status         status    `db:"status" json:""`
	Company        string    `db:"company" json:"company"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type Group struct {
	ID        int32     `db:"id" json:"id"`
	Name      string    `db:"group_name" json:"group_name"`
	Desc      string    `db:"group_bio" json:"description"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatorID int32     `db:"creator_id" json:"creator_id"`
}

type MemberShip struct {
	ID       int32     `db:"id" json:"id"`
	UserID   int32     `db:"user_id" json:"user_id"`
	GroupID  int32     `db:"group_id" json:"group_id"`
	JoinedAt time.Time `db:"joined_at" json:"joined_at"`
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
	user, err := s.driver.GetUserByUsername(ctx, username)
	if err != nil {
		return false
	}
	return user.Username == username
}

func (s *Store) GetUserByID(ctx context.Context, u_id int32) (*User, error) {
	return s.driver.GetUserByID(ctx, u_id)
}

func (s *Store) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return s.driver.GetUserByUsername(ctx, username)
}

func (s *Store) AddUserToCache(username string) {
	s.users.Store(username, true)
}

func (s *Store) GetAllGroups(ctx context.Context) []Group {
	return s.driver.GetGroups(ctx)
}

func (s *Store) JoinGroup(ctx context.Context, g_id, u_id int32) error {
	return s.driver.JoinGroup(ctx, g_id, u_id)
}
