package postgresql

import (
	"context"
	"errors"
	"fmt"

	"tdevs/store"
)

func (db *PostgresDB) CreateUser(ctx context.Context, user *store.User) error {
	db.db.Exec("INSERT INTO USERS (username, hashed_password, company, created_at, updated_at) VALUES ($1, $2,$3,$4,$5)",
		user.Username, user.Hashed_Password, user.Company, user.Created_At, user.Updated_At)
	return nil
}

func (db *PostgresDB) GetUser(ctx context.Context, username string) (*store.User, error) {

	users := []store.User{}

	err := db.db.Select(&users, "SELECT * from users where username=$1 LIMIT 1", username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user does not exist")
	}

	return &users[0], nil
}
