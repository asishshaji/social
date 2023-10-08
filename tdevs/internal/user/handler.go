package user

import (
	"net/http"
	"tdevs/data"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	uS userService
}

func NewUserHandler(uS userService) userHandler {
	return userHandler{
		uS: uS,
	}
}

func (uH userHandler) RegisterUser(c echo.Context) error {
	uDto := data.UserDTO{}
	if err := c.Bind(&uDto); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	u, err := uH.uS.CreateUser(c, uDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"username": u.Username,
	})
}

func (uH userHandler) Login(c echo.Context) error {
	lDto := data.UserLoginDTO{}
	if err := c.Bind(&lDto); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	res, err := uH.uS.Login(c, lDto)
	if err != nil {
		return c.JSON(res.StatusCode, err.Error())
	}

	return c.JSON(res.StatusCode, res)

}
