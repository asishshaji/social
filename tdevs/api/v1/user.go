package v1

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

const CONTEXT_KEY = "token"

func jwt_custom_middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, "you are not authorized")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, "you are not authorized")
		}
		u_id := int32(claims["u_id"].(float64))
		c.Set("u_id", u_id)

		return next(c)
	}
}

func (s *APIV1Service) registerUserRoutes(g *echo.Group) {
	jwt_config := echojwt.Config{
		ContextKey: CONTEXT_KEY,
		SigningKey: []byte(s.Profile.JWT_Secret),
	}
	g.Use(echojwt.WithConfig(jwt_config))
	g.Use(jwt_custom_middleware)
	g.GET("/user/me", s.GetCurrentUser)
	g.GET("/user/group/all", s.GetGroups)
	g.POST("/user/group/join/:id", s.JoinGroup)
}

func (api *APIV1Service) JoinGroup(c echo.Context) error {
	g_id := c.Param("id")
	i, err := strconv.ParseInt(g_id, 10, 32)
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	g_id_i32 := int32(i)
	ctx := c.Request().Context()
	u_id := c.Get("u_id").(int32)

	err = api.Store.JoinGroup(ctx, g_id_i32, u_id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]string{})
}

func (api *APIV1Service) GetCurrentUser(c echo.Context) error {
	ctx := c.Request().Context()
	u_id := c.Get("u_id").(int32)
	user, err := api.Store.GetUserByID(ctx, u_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (api *APIV1Service) GetGroups(c echo.Context) error {
	groups := api.Store.GetAllGroups(c.Request().Context())
	return c.JSON(http.StatusAccepted, groups)
}
