package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Logger.Fitness/backend/controllers"
	"Logger.Fitness/backend/db"
)

type ContextParams struct {
	DbClient *db.Client
}

func Run(dbClient *db.Client, port string) {
	contextParams := ContextParams{DbClient: dbClient}

	e := echo.New()
	e.Use(
		createContext(contextParams),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.CORS(),
	)

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${method} ${status} ${uri} ${user_agent}\n",
	// }))

	e = initRoutes(e)

	e.Logger.Fatal(e.Start(port))
}

func initRoutes(e *echo.Echo) *echo.Echo {

	e.POST("/auth/verify_me", controllers.VerifyMe)
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	return e
}

// ContextObjects attaches backend clients to the API context to be
//	used inside controllers
func createContext(contextParams ContextParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", contextParams.DbClient)
			return next(c)
		}
	}
}
