package v1

import (
	"net/http"
	"tdevs/store"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	Password string `json:"password"`
	Company  string `json:"company"`
}

func (s *APIV1Service) registerAuthRoutes(group *echo.Group) {
	group.POST("/auth/signup", s.signup)
}

func (s *APIV1Service) signup(c echo.Context) error {
	ctx := c.Request().Context()
	signUp := SignUp{}
	if err := c.Bind(&signUp); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate password hash").SetInternal(err)
	}

	user, err := s.Store.CreateUser(ctx, &store.User{
		Username: signUp.Company,
		Company:  signUp.Company,
		Password: string(passwordHash),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user").SetInternal(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"username": user.Username,
	})
}
