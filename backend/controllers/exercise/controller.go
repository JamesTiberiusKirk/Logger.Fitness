package exercise

import (
	"net/http"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const path = "/exercises"

// DatabaseInterface ...
type DatabaseInterface interface {
	InsertNewExercise(newExercise types.Exercise) (types.Exercise, error)
	GetUserExercises(userID primitive.ObjectID) ([]types.Exercise, error)
	GetUserExercisesByWorkout(userID, workoutID primitive.ObjectID) ([]types.Exercise, error)
	EditExercise(exercise types.Exercise) error
	DeleteExerciseByID(userID, exerciseID primitive.ObjectID) error
}

// ExerciseController ...
type ExerciseController struct {
	database       DatabaseInterface
	authMiddleware echo.MiddlewareFunc
}

// NewExerciseController ...
func NewExerciseController(database DatabaseInterface, authMiddleware echo.MiddlewareFunc) ExerciseController {
	return ExerciseController{
		database:       database,
		authMiddleware: authMiddleware,
	}
}

// Init ...
func (ctrl *ExerciseController) Init(g *echo.Group) {
	group := g.Group(path)

	group.GET("", ctrl.GetExercise, ctrl.authMiddleware)
	group.POST("", ctrl.PostExercise, ctrl.authMiddleware)
	group.PUT("", ctrl.PutExercise, ctrl.authMiddleware)
	group.DELETE("", ctrl.DeleteExercise, ctrl.authMiddleware)

}

// GetExercise GET endpoint.
// @param workout_id - optional param
func (ctrl *ExerciseController) GetExercise(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	// IF a workout id is provided it will return exercises in that workout
	workoitID := c.QueryParam("workout_id")
	if workoitID == "" {
		workoitIDObj, err := primitive.ObjectIDFromHex(workoitID)
		if err != nil {
			log.Error(err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		log.Print(workoitID)
		log.Print(workoitIDObj.Hex())

		results, err := db.GetUserExercisesByWorkout(userClaim.ID, workoitIDObj)
		if err != nil {
			log.Error(err.Error())
			return c.String(http.StatusInternalServerError, res.DatabseError)
		}
		return c.JSON(http.StatusOK, results)
	}

	results, err := db.GetUserExercises(userClaim.ID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	if len(results) == 0 {
		return c.JSON(http.StatusOK, []types.Exercise{})
	}

	return c.JSON(http.StatusOK, results)
}

// PostExercise POST endpoint.
// Starts new exercise.
// @body - expected type types.Exercise
func (ctrl *ExerciseController) PostExercise(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	var newExercise types.Exercise
	if err := c.Bind(&newExercise); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newExercise.UserID = userClaim.ID
	exercise, err := db.InsertNewExercise(newExercise)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, exercise)

}

// PutExercise PUT endpoint.
// Deletes exercise.
// TODO: BUG: this will accept either data type ignoring the set type
func (ctrl *ExerciseController) PutExercise(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	var userUpdate types.Exercise
	err := c.Bind(&userUpdate)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}
	userUpdate.UserID = userClaim.ID
	err = db.EditExercise(userUpdate)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// DeleteExercise DELETE endpoint.
// Deletes exercise.
// @param - exercise_id
func (ctrl *ExerciseController) DeleteExercise(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseID, err := primitive.ObjectIDFromHex(c.QueryParam("exercise_id"))
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	err = db.DeleteExerciseByID(userClaim.ID, exerciseID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}
