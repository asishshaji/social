package postgresql

import (
	"context"
	"errors"
	"fmt"

	"tdevs/store"
)

func (postgres *PostgresDB) CreateUser(ctx context.Context, user *store.User) error {
	_, err := postgres.db.Exec("INSERT INTO USERS (username, HashedPassword, company, CreatedAt, UpdatedAt) VALUES ($1, $2,$3,$4,$5)",
		user.Username, user.HashedPassword, user.Company, user.CreatedAt, user.UpdatedAt)
	return err
}

func (postgres *PostgresDB) GetUser(ctx context.Context, username string) (*store.User, error) {
	users := []store.User{}

	err := postgres.db.Select(&users, "SELECT * FROM WHERE username=$1 LIMIT 1", username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user does not exist")
	}

	return &users[0], nil
}
