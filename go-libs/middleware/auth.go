package lfMiddleware

import (
	"Logger.Fitness/go-libs/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Available auth roles
const (
	AdminRole = "admin"
	UserRole  = "user"
)

// Auth returns echo middleware function which validates jwt on time
//	and role specified. This also revalidated the jwt token.
func Auth(role string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:x-access-token",
		Validator: func(jwt string, c echo.Context) (bool, error) {
			claim, err := auth.ValidateJWTToken(jwt)
			if err != nil {
				return false, err
			}

			err = claim.Valid()
			if err != nil {
				return false, err
			}
			if _, found := claim.Roles[role]; !found {
				return false, err
			}

			newJwt, err := auth.GenerateJWTFromClaim(*claim)
			if err != nil {
				return false, err
			}

			c.Response().Header().Add("x-access-token", newJwt)
			c.Set("user", claim)
			return true, err
		},
	})
}
