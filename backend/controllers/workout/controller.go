package workout

import (
	"net/http"
	"strconv"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const path = "/workouts"

// DatabaseInterface ...
type DatabaseInterface interface {
}

// WorkoutController ...
type WorkoutController struct {
	database       DatabaseInterface
	authMiddleware echo.MiddlewareFunc
}

// NewWorkoutController ...
func NewWorkoutController(database DatabaseInterface, authMiddleware echo.MiddlewareFunc) WorkoutController {
	return WorkoutController{
		database:       database,
		authMiddleware: authMiddleware,
	}
}

// Init ...
func (ctrl *WorkoutController) Init(g *echo.Group) {
	group := g.Group(path)

	group.GET("", ctrl.GetWorkouts)
	group.DELETE("", ctrl.DeleteWorkout)

}

// WorkoutDTO - workout & their exercises Data Transfer Object
type WorkoutDTO struct {
	Workout   types.Workout    `json:"workout"`
	Exercises []types.Exercise `json:"exercises"`
}

// StartNewWorkout POST endpoint.
// Starts a new workout if there are no other
// 	workouts active.
// @body - types.Workout (only send start time, title and notes, the rest will
//	be overwritten)
// returns newly created record
// TODO: validate start_time
func (ctrl *WorkoutController) StartNewWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	activeWorkout, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		isResultEmpty := err.Error() == "mongo: no documents in result"
		if !isResultEmpty {
			log.Error(err.Error())
			return c.String(http.StatusInternalServerError, res.DatabseError)
		}
	}

	var emptyWorkout types.Workout
	if activeWorkout != emptyWorkout {
		return c.String(http.StatusBadRequest, "Need to stop previous workout first")
	}

	var newWorkout types.Workout
	err = c.Bind(&newWorkout)
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	newWorkout.UserID = userClaim.ID
	newWorkout.ID = primitive.NewObjectID()
	newWorkout.EndTime = -1

	err = db.InsertNewWorkout(newWorkout)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)

	}

	return c.JSON(http.StatusOK, newWorkout)
}

// StopWorkout GET endpoint.
// Stops the active workout.
// This also leaves an end time time stamp.
// @param end_time - Unix timestamp
// returns modified record
// TODO: validate end_time
func (ctrl *WorkoutController) StopWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)
	endTimeStamp, err := strconv.ParseInt(c.QueryParam("end_time"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	activeWorkout, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	activeWorkout.EndTime = endTimeStamp
	err = db.EditWorkout(activeWorkout)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, activeWorkout)
}

// GetWorkouts GET endpoint..
// Gets all user workouts.
// Return
// @body - []WorkoutDTO
func (ctrl *WorkoutController) GetWorkouts(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	workouts, err := db.GetUserWorkouts(userClaim.ID)

	if err != nil {
		isResultEmpty := err.Error() == "mongo: no documents in result"
		if !isResultEmpty {
			log.Error(err.Error())
			return c.String(http.StatusInternalServerError, res.DatabseError)
		}
		log.Error(err.Error())
	}

	var results []WorkoutDTO
	for i := 0; i < len(workouts); i++ {
		exercises, err := db.GetUserExercisesByWorkout(userClaim.ID, workouts[i].ID)
		if err != nil {
			log.Error(err.Error())
			return c.String(http.StatusInternalServerError, res.DatabseError)
		}

		item := WorkoutDTO{
			Workout:   workouts[i],
			Exercises: exercises,
		}
		results = append(results, item)
	}

	return c.JSON(http.StatusOK, results)
}

// GetActiveWorkout GET endpoint.
// Gets the one active workout a user has
func (ctrl *WorkoutController) GetActiveWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	result, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, result)
}

// DeleteWorkout DELETE endpoint.
// Deletes a workout
// @param workout_id - id of workout to delete
func (ctrl *WorkoutController) DeleteWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)
	workoutID, err := primitive.ObjectIDFromHex(c.QueryParam("workout_id"))
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err = db.DeleteWorkout(workoutID, userClaim.ID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	err = db.DeleteExerciseByWorkoutID(userClaim.ID, workoutID)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// EditWorkout PUT endpoint.
// Re-writes a workout.
// @body update - types.Workout
func (ctrl *WorkoutController) EditWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	var update types.Workout
	err := c.Bind(&update)
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	update.UserID = userClaim.ID
	err = db.EditWorkout(update)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}
