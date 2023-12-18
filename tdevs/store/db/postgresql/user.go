package postgresql

import (
	"context"
	"errors"
	"fmt"

	"tdevs/store"
)

func (postgres *PostgresDB) CreateUser(ctx context.Context, user *store.User) error {
	_, err := postgres.db.Exec("INSERT INTO users (username, hashed_password, company, created_at, updated_at) VALUES ($1, $2,$3,$4,$5)",
		user.Username, user.HashedPassword, user.Company, user.CreatedAt, user.UpdatedAt)
	return err
}

func (postgres *PostgresDB) GetUserByUsername(ctx context.Context, username string) (*store.User, error) {
	users := []store.User{}

	err := postgres.db.Select(&users, "SELECT * FROM users WHERE username=$1 LIMIT 1", username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user does not exist")
	}

	return &users[0], nil
}

func (postgres *PostgresDB) GetUserByID(ctx context.Context, u_id int32) (*store.User, error) {
	users := []store.User{}

	err := postgres.db.Select(&users, "SELECT * FROM users WHERE id=$1 LIMIT 1", u_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user does not exist")
	}

	return &users[0], nil
}

func (postgres *PostgresDB) GetGroupsForUser(ctx context.Context, user_id int32) []store.Group {
	groups := []store.Group{}

	// err := postgres.db.Select(&groups, "SELECT * FROM groups WHERE =$1 LIMIT 1", username)

	return groups
}
