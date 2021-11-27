package controllers

import (
	"net/http"
	"strconv"
	"time"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Workouts */

// GetWorkouts GET endpoint..
// Gets all user workouts.
func GetWorkouts(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	result, err := db.GetUserWorkouts(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, result)
}

// GetActiveWorkout GET endpoint.
// Gets the one active workout a user has
func GetActiveWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	result, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, result)
}

// StartNewWorkout POST endpoint.
// Starts a new workout if there are no other
// 	workouts active.
// @body - start_time unix time
func StartNewWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	activeWorkout, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}
	if activeWorkout != nil {
		return c.String(http.StatusBadRequest, "Need to stop previous workout first")
	}

	var newWorkout types.Workout
	err = c.Bind(&newWorkout)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	newWorkout.UserID = userClaim.ID
	newWorkout.ID = primitive.NewObjectID()
	newWorkout.EndTime = nil

	startTime := c.QueryParam("start_time")
	startTimeInt, err := strconv.ParseInt(startTime, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, res.CannotParseString)
	}
	newWorkout.StartTime = time.Unix(startTimeInt, 0)

	err = db.InsertNewWorkout(newWorkout)

	return c.NoContent(http.StatusOK)
}

// StopWorkout GET endpoint.
// Stops the active workout.
// This also leaves an end time time stamp.
// @param end_time - RFC3339 timestamp
func StopWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)
	endTimeStamp, err := time.Parse(time.RFC3339, c.QueryParam("end_time"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	activeWorkout, err := db.GetUserAcitveWorkout(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	activeWorkout.EndTime = &endTimeStamp
	err = db.EditWorkout(*activeWorkout)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.String(http.StatusOK, "Workout stopped")
}

// DeleteWorkout DELETE endpoint.
// Deletes a workout
// @param workout_id - id of workout to delete
func DeleteWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)
	workoutID, err := primitive.ObjectIDFromHex(c.QueryParam("workout_id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err = db.DeleteWorkout(workoutID, userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// EditWorkout PUT endpoint.
// Re-writes a workout.
// NOTE: This will even re-write blank fields.
// @body update - types.Workout
func EditWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	var update types.Workout
	err := c.Bind(&update)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	update.UserID = userClaim.ID
	err = db.EditWorkout(update)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}
