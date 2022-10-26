package server

import (
	"Logger.Fitness/backend/controllers/auth"
	"Logger.Fitness/backend/controllers/exercise"
	"Logger.Fitness/backend/controllers/exercise_type"
	"Logger.Fitness/backend/controllers/workout"
	"Logger.Fitness/backend/db"
	lfMiddleware "Logger.Fitness/go-libs/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

// ContextParams is for the context passed to the controllers
type ContextParams struct {
	DbClient *db.DbClient
}

// Run runs the http server
func Run(dbClient *db.DbClient, port string, googleOauthConfig *oauth2.Config) {
	contextParams := ContextParams{DbClient: dbClient}

	e := echo.New()
	e.Use(
		createContext(contextParams),
		middleware.Logger(),
		middleware.CORS(),
	)

	userAuthMiddleware := lfMiddleware.Auth(lfMiddleware.UserRole)

	g := e.Group("/api/v3")

	authGroup := auth.NewAuthController(dbClient, googleOauthConfig)
	authGroup.Init(g)

	exercioseGroup := exercise.NewExerciseController(dbClient, userAuthMiddleware)
	exercioseGroup.Init(g)

	exerciseTypeGroup := exercise_type.NewExerciseTypeController(dbClient, userAuthMiddleware)
	exerciseTypeGroup.Init(g)

	// Need to create the workout controller
	workoutGroup := workout.NewController(dbClient, userAuthMiddleware)
	workoutGroup.Init(g)

	e.Logger.Fatal(e.Start(port))
}

// ContextObjects attaches backend clients to the API context to be
//
//	used inside controllers
func createContext(contextParams ContextParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", contextParams.DbClient)
			return next(c)
		}
	}
}
