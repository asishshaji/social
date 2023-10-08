package user

import (
	"errors"
	"fmt"
	"tdevs/data"
	"tdevs/internal/cache"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo  IUserRepo
	cache cache.ICache
}

func NewUserService(repo IUserRepo, cache cache.ICache) userService {
	return userService{
		repo:  repo,
		cache: cache,
	}
}

func (s userService) CreateUser(c echo.Context, dto data.UserDTO) (data.User, error) {
	retryCount := 5
	exists := true

	u := data.User{}

	var err error

	for exists {
		u.Username = petname.Generate(2, "_")
		exists, err = s.repo.CheckUserNameExists(c.Request().Context(), u.Username)
		if err != nil {
			return u, err
		}
		if retryCount <= 0 {
			c.Logger().Info("error creating uname")
			return u, errors.New("error creating uname")
		}
		retryCount--
	}
	u.Company = dto.Company

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 5)
	if err != nil {
		return u, fmt.Errorf("error generating password hash: %s", err)
	}
	u.Password = string(hashedPassword)

	err = s.repo.InsertUser(c.Request().Context(), u)
	if err != nil {
		return u, fmt.Errorf("error creating u %s", err)
	}

	return u, nil

}

func (s userService) Login(c echo.Context, uDto data.UserLoginDTO) (data.UserLoginResponse, error) {
	res := data.UserLoginResponse{}
	hashedPassword, err := s.repo.GetUserPassword(c.Request().Context(), uDto.Username)
	if err != nil {
		res.StatusCode = 404
		return res, fmt.Errorf("error getting user : %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(uDto.Password))
	if err != nil {
		res.StatusCode = 403
		return res, fmt.Errorf("invalid password : %s", err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["user"] = uDto.Username
	claims["authorized"] = true

	tokenStr, err := token.SignedString([]byte("ASD"))
	if err != nil {
		res.StatusCode = 500
		return res, fmt.Errorf("error generating access token : %s", err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["exp"] = time.Now().Add(10 * time.Hour * 24).Unix()
	refreshClaims["user"] = uDto.Username
	refreshClaims["authorized"] = true

	refreshTokenStr, err := refreshToken.SignedString([]byte("ASD"))
	if err != nil {
		res.StatusCode = 500
		return res, fmt.Errorf("error generating refresh token : %s", err)
	}

	res.AccessToken = tokenStr
	res.RefreshToken = refreshTokenStr
	res.StatusCode = 200

	// TODO persistence for session
	// https://redis.com/blog/json-web-tokens-jwt-are-dangerous-for-user-sessions/
	s.cache.SetAuthenticatedUser(c.Request().Context(), uDto.Username)
	return res, nil
}
