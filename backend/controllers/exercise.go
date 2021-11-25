package controllers

import (
	"github.com/labstack/echo/v4"
)

/* Exercise */

// TODO: Need to finish all operations on exercises

// NewExercise POST endpoint.
// Starts new exercise.
func NewExercise(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// DeleteExercise DELETE endpoint.
// Deletes exercise.
func DeleteExercise(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

/* Sets in exercises */

// AddSetInExercise POST endpoint.
// Added new set in exercise.
func AddSetInExercise(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// NOTE: Maybe the following two can be left out and
// 	only update either the exercise or the entire workout.

// EditSetInExercise endpoint.
// Edit an individual set in an exercise.
// NOTE: Precision specified by array iterator.
func EditSetInExercise(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}

// DeleteSetInExercise DELETE endpoint.
// Deletes set in exercise.
func DeleteSetInExercise(c *echo.Context) error {
	// db := c.Get("db").(*db.DbClient)
	// userClaim := c.Get("user").(*types.JwtClaim)
	return nil
}
