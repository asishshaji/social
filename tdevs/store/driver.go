package store

import "context"

type IDriver interface {
	CreateUser(context.Context, *User) (*User, error)
}
