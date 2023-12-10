package v1

import (
	"net/http"
	"tdevs/store"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	Password string `json:"password"`
	Company  string `json:"company"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refersh_token"`
}

func (api *APIV1Service) registerAuthRoutes(group *echo.Group) {
	group.POST("/auth/signup", api.signup)
	group.POST("/auth/signin", api.signin)
}

func (api *APIV1Service) signup(c echo.Context) error {
	ctx := c.Request().Context()

	signUp := SignUp{}
	if err := c.Bind(&signUp); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate password hash").SetInternal(err)
	}

	exists := true
	retryCount := 5
	var username string
	for exists {
		username = petname.Generate(2, "_")
		exists = api.Store.CheckUserExists(ctx, username)
		if retryCount <= 0 {
			break
		}
		retryCount--
	}

	if retryCount <= 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate a unique username")
	}

	user, err := api.Store.CreateUser(ctx, &store.User{
		Username:       signUp.Company,
		Company:        signUp.Company,
		HashedPassword: string(passwordHash),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user").SetInternal(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"username": user.Username,
	})
}

func (api *APIV1Service) signin(c echo.Context) error {
	ctx := c.Request().Context()

	login := Login{}
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	user, err := api.Store.GetUser(ctx, login.Username)
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(login.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	accessToken, err := generateToken(user.Username, api.Profile.JWT_Secret, 10*time.Minute)
	if err != nil {
		return echo.ErrInternalServerError
	}
	refreshToken, err := generateToken(user.Username, api.Profile.JWT_Secret, 10*time.Hour*24)

	if err != nil {
		return echo.ErrInternalServerError
	}

	api.Store.AddUserToCache(user.Username)

	return c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}

func generateToken(username, secret string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
