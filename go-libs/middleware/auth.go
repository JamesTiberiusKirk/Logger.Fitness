package lfMiddleware

import (
	lib "Logger.Fitness/go-libs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Available auth roles
const (
	AdminRole = "admin"
	UserRole  = "user"
)

// Auth returns echo middleware function which validates jwt on time
//	and role specified
// TODO: need to revalidate the jwt token and reinsert it into header
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

			newClaims, err := lib.GenerateJWTFromClaim(*claims)
			if err != nil {
				return false, err
			}

			c.Set("user", newClaims)
			return true, err
		},
	})
}
