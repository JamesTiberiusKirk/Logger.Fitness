package workout

import (
	"net/http"
	"strconv"

	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const path = "/workouts"

// DatabaseInterface ...
type DatabaseInterface interface {
	GetUserAcitveWorkout(userID primitive.ObjectID) (types.Workout, error)
	InsertNewWorkout(newWorkout types.Workout) error
	GetUserWorkouts(userID primitive.ObjectID) ([]types.Workout, error)
	GetUserExercisesByWorkout(userID, workoutID primitive.ObjectID) ([]types.Exercise, error)
	DeleteWorkout(workoutID, userID primitive.ObjectID) error
	EditWorkout(workout types.Workout) error
	DeleteExerciseByWorkoutID(userID, workoutID primitive.ObjectID) error
}

// Controller ...
type Controller struct {
	database       DatabaseInterface
	authMiddleware echo.MiddlewareFunc
}

// NewController ...
func NewController(database DatabaseInterface, authMiddleware echo.MiddlewareFunc) Controller {
	return Controller{
		database:       database,
		authMiddleware: authMiddleware,
	}
}

// Init ...
func (ctrl *Controller) Init(g *echo.Group) {
	group := g.Group(path)

	group.PUT("", ctrl.EditWorkout, ctrl.authMiddleware)
	group.GET("", ctrl.GetWorkouts, ctrl.authMiddleware)
	group.DELETE("", ctrl.DeleteWorkout, ctrl.authMiddleware)
}

// DTO - workout & their exercises Data Transfer Object
type DTO struct {
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
func (ctrl *Controller) StartNewWorkout(c echo.Context) error {
	db := ctrl.database
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
func (ctrl *Controller) StopWorkout(c echo.Context) error {
	db := ctrl.database
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
// @body - []DTO
func (ctrl *Controller) GetWorkouts(c echo.Context) error {
	db := ctrl.database
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

	var results []DTO
	for i := 0; i < len(workouts); i++ {
		exercises, err := db.GetUserExercisesByWorkout(userClaim.ID, workouts[i].ID)
		if err != nil {
			log.Error(err.Error())
			return c.String(http.StatusInternalServerError, res.DatabseError)
		}

		item := DTO{
			Workout:   workouts[i],
			Exercises: exercises,
		}
		results = append(results, item)
	}

	// FIX, if 0 items, returns null
	if len(workouts) == 0 {
		return c.JSON(http.StatusOK, []DTO{})
	}

	return c.JSON(http.StatusOK, results)
}

// GetActiveWorkout GET endpoint.
// Gets the one active workout a user has
func (ctrl *Controller) GetActiveWorkout(c echo.Context) error {
	db := ctrl.database
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
func (ctrl *Controller) DeleteWorkout(c echo.Context) error {
	db := ctrl.database
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
func (ctrl *Controller) EditWorkout(c echo.Context) error {
	db := ctrl.database
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
