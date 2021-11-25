package controllers

import (
	"github.com/labstack/echo/v4"
)

/* Workouts */

// GetWorkouts GET endpoint..
// Gets all user workouts.
func GetWorkouts(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)

	return nil
}

// GetActiveWorkout GET endpoint.
// Gets the one active workout a user has
func GetActiveWorkout(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// StartNewWorkout POST endpoint.
// Starts a new workout if there are no other
// 	workouts active.
func StartNewWorkout(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// StopWorkout GET endpoint.
// Stops the active workout.
// This also leaves an end time time stamp.
func StopWorkout(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// DeleteWorkout DELETE endpoint.
// Deletes a workout
func DeleteWorkout(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// EditWorkout PUT endpoint.
// Re-writes a workout.
// NOTE: This will even re-write blank fields.
func EditWorkout(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}
