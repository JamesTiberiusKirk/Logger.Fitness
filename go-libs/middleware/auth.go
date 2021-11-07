package lfMiddleware

import (
	lib "Logger.Fitness/go-libs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	ADMIN_ROLE = "admin"
	USER_ROLE  = "user"
)

// Auth returns echo middleware function which validates jwt on time
//	and role specified
func Auth(role string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:jwt-token",
		Validator: func(jwt string, c echo.Context) (bool, error) {
			claims, err := lib.ValidateJwtToken(jwt)
			if err != nil {
				return false, err
			}

			err = claims.Valid()
			if err != nil {
				return false, err
			}
			if _, found := claims.Roles[role]; !found {
				return false, err
			}

			c.Set("user", claims)
			return true, err
		},
	})
}
