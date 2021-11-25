package controllers

import (
	"net/http"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Workouts */

// TODO: test

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

	result, err := db.GetUserAcitveWorkouts(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.JSON(http.StatusOK, result)
}

// StartNewWorkout POST endpoint.
// Starts a new workout if there are no other
// 	workouts active.
func StartNewWorkout(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	activeWorkout, err := db.GetUserAcitveWorkouts(userClaim.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}
	if len(activeWorkout) != 0 {
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

	err = db.InsertNewWorkout(newWorkout)

	return c.NoContent(http.StatusOK)
}

// StopWorkout GET endpoint.
// Stops the active workout.
// This also leaves an end time time stamp.
func StopWorkout(c echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// DeleteWorkout DELETE endpoint.
// Deletes a workout
func DeleteWorkout(c echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// EditWorkout PUT endpoint.
// Re-writes a workout.
// NOTE: This will even re-write blank fields.
func EditWorkout(c echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}
