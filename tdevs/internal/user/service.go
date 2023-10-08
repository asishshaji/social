package user

import (
	"errors"
	"fmt"
	"os"
	"tdevs/data"
	"tdevs/internal/cache"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo      IUserRepo
	JWTSecret string
	cache     cache.ICache
	usernames chan string
}

func NewUserService(repo IUserRepo, cache cache.ICache) userService {

	uS := userService{
		repo:      repo,
		cache:     cache,
		JWTSecret: os.Getenv("JWT_SECRET"),
		usernames: make(chan string, 100),
	}

	go uS.generateRandomUsernames()

	return uS
}

func (uS userService) generateRandomUsernames() {
	for {
		if len(uS.usernames) < 30 {
			uS.usernames <- petname.Generate(2, "_")
		} else {
			time.Sleep(time.Second * 2)
		}
	}
}

func (uS userService) CreateUser(c echo.Context, dto data.UserDTO) (data.User, error) {
	retryCount := 5
	exists := true

	u := data.User{}

	var err error

	for exists {
		// run a goroutine that pushes names to a channel,
		// sort of prepopulate random names
		u.Username = <-uS.usernames
		exists, err = uS.repo.CheckUserNameExists(c.Request().Context(), u.Username)
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

	err = uS.repo.InsertUser(c.Request().Context(), u)
	if err != nil {
		return u, fmt.Errorf("error creating u %s", err)
	}

	return u, nil

}

func (uS userService) Login(c echo.Context, uDto data.UserLoginDTO) (data.UserLoginResponse, error) {
	res := data.UserLoginResponse{}
	hashedPassword, err := uS.repo.GetUserPassword(c.Request().Context(), uDto.Username)
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

	tokenStr, err := token.SignedString([]byte(uS.JWTSecret))
	if err != nil {
		res.StatusCode = 500
		return res, fmt.Errorf("error generating access token : %s", err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["exp"] = time.Now().Add(10 * time.Hour * 24).Unix()
	refreshClaims["user"] = uDto.Username
	refreshClaims["authorized"] = true

	refreshTokenStr, err := refreshToken.SignedString([]byte(uS.JWTSecret))
	if err != nil {
		res.StatusCode = 500
		return res, fmt.Errorf("error generating refresh token : %s", err)
	}

	res.AccessToken = tokenStr
	res.RefreshToken = refreshTokenStr
	res.StatusCode = 200

	uS.cache.SetAuthenticatedUser(c.Request().Context(), uDto.Username)
	return res, nil
}
