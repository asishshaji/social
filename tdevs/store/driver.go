package store

import "context"

type IDriver interface {
	CreateUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, error)
}
