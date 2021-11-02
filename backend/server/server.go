package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Logger.Fitness/backend/controllers"
	"Logger.Fitness/backend/db"
	lfMiddleware "Logger.Fitness/go-libs/middleware"
)

type ContextParams struct {
	DbClient *db.DbClient
}

func Run(dbClient *db.DbClient, port string) {
	contextParams := ContextParams{DbClient: dbClient}

	e := echo.New()
	e.Use(
		createContext(contextParams),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.CORS(),
	)

	e = initRoutes(e)

	e.Logger.Fatal(e.Start(port))
}

func initRoutes(e *echo.Echo) *echo.Echo {

	e.POST("/auth/verify_me", controllers.VerifyMe)
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	e.GET("/helloworld", controllers.HelloWorld, lfMiddleware.Auth(lfMiddleware.USER_ROLE))

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
