package store

import "context"

type IDriver interface {
	CreateUser(context.Context, *User) error
	GetUserByUsername(context.Context, string) (*User, error)
	GetUserByID(context.Context, int32) (*User, error)
	GetGroups(ctx context.Context) []Group
	GetGroupsForUser(ctx context.Context, user_id int32) []Group
	JoinGroup(context.Context, int32, int32) error
}
