package server

import (
	"Logger.Fitness/backend/controllers"
	"Logger.Fitness/backend/db"
	lfMiddleware "Logger.Fitness/go-libs/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ContextParams is for the context passed to the controllers
type ContextParams struct {
	DbClient *db.DbClient
}

// Run runs the http server
func Run(dbClient *db.DbClient, port string) {
	contextParams := ContextParams{DbClient: dbClient}

	e := echo.New()
	e.Use(
		createContext(contextParams),
		middleware.Logger(),
		middleware.CORS(),
	)

	e = initRoutes("/api/v2", e)

	e.Logger.Fatal(e.Start(port))
}

func initRoutes(prefix string, e *echo.Echo) *echo.Echo {
	userAuth := lfMiddleware.Auth(lfMiddleware.UserRole)

	e.GET(prefix+"/hw", controllers.HelloWorld)

	e.POST(prefix+"/auth/verify_me", controllers.VerifyMe)
	e.POST(prefix+"/auth/register", controllers.Register)
	e.POST(prefix+"/auth/login", controllers.Login)

	e.POST(prefix+"/exercise_type", controllers.NewExerciseType, userAuth)
	e.GET(prefix+"/exercise_type", controllers.GetExerciseTypes, userAuth)
	e.PUT(prefix+"/exercise_type", controllers.EditExerciseTypes, userAuth)
	e.DELETE(prefix+"/exercise_type", controllers.DeleteExerciseType, userAuth)

	e.POST(prefix+"/workouts/start", controllers.StartNewWorkout, userAuth)
	e.POST(prefix+"/workouts/stop", controllers.StopWorkout, userAuth)
	e.GET(prefix+"/workouts/active", controllers.GetActiveWorkout, userAuth)
	e.GET(prefix+"/workouts", controllers.GetWorkouts, userAuth)
	e.PUT(prefix+"/workouts", controllers.EditWorkout, userAuth)
	e.DELETE(prefix+"/workouts", controllers.DeleteWorkout, userAuth)

	e.GET(prefix+"/exercises", controllers.GetExercise, userAuth)
	e.POST(prefix+"/exercises", controllers.PostExercise, userAuth)
	e.PUT(prefix+"/exercises", controllers.PutExercise, userAuth)
	e.DELETE(prefix+"/exercises", controllers.DeleteExercise, userAuth)

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
