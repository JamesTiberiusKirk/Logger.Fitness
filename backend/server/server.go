package server

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"Logger.Fitness/backend/controllers"
	"Logger.Fitness/backend/db"
	lfMiddleware "Logger.Fitness/go-libs/middleware"
)

// ContextParams is for the context passed to the controllers
type ContextParams struct {
	DbClient *db.DbClient
}

// Run runs the http server
// @title Logger.Fitness API
// @version 0.0
// @description This is a sample server Petstore server.
// @host petstore.swagger.io
// @BasePath /
func Run(dbClient *db.DbClient, port string) {
	contextParams := ContextParams{DbClient: dbClient}

	e := echo.New()
	e.Use(
		createContext(contextParams),
		middleware.Logger(),
		middleware.CORS(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				if strings.Contains(c.Request().URL.Path, "swagger") {
					return true
				}
				return false
			},
		}),
	)

	e = initRoutes(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(port))
}

func initRoutes(e *echo.Echo) *echo.Echo {

	e.POST("/auth/verify_me", controllers.VerifyMe)
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	e.GET("/helloworld", controllers.HelloWorld, lfMiddleware.Auth(lfMiddleware.UserRole))

	e.POST("/extp", controllers.NewExerciseType, lfMiddleware.Auth(lfMiddleware.UserRole))
	e.GET("/extp", controllers.GetExerciseTypes, lfMiddleware.Auth(lfMiddleware.UserRole))
	e.PUT("/extp", controllers.EditExerciseTypes, lfMiddleware.Auth(lfMiddleware.UserRole))

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
