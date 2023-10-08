package user

import (
	"context"
	"fmt"
	"log"
	"tdevs/data"

	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	InsertUser(context.Context, data.User) error
	CheckUserNameExists(context.Context, string) (bool, error)
	GetUserPassword(context.Context, string) (string, error)
}

type UserRepo struct {
	l  *log.Logger
	db *sqlx.DB
}

func NewUserRepo(l *log.Logger, db *sqlx.DB) IUserRepo {
	return &UserRepo{l: l, db: db}
}

func (r UserRepo) CheckUserNameExists(c context.Context, username string) (bool, error) {
	var exists []bool
	err := r.db.Select(&exists, "SELECT EXISTS(SELECT 1 from users where username=$1);", username)
	if err != nil {
		r.l.Println(err)
		return false, err
	}

	return exists[0], nil
}

func (r UserRepo) InsertUser(ctx context.Context, u data.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, company) VALUES ($1, $2, $3);", u.Username, u.Password, u.Company)
	return err
}

func (r UserRepo) GetUserPassword(c context.Context, username string) (string, error) {
	var u []string
	err := r.db.Select(&u, "SELECT password from users where username=$1", username)
	if err != nil {
		r.l.Println(err)
		return "", err
	}

	if len(u) < 1 {
		r.l.Println("username doesnot exist")
		return "", fmt.Errorf("username %s doesnot exist", username)
	}

	return u[0], nil
}
