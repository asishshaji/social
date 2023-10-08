package data

import "time"

type User struct {
	Username  string
	Password  string
	Company   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDTO struct {
	Password string `json:"password"`
	Company  string `json:"company"`
}

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	StatusCode   int    `json:"-"`
}
