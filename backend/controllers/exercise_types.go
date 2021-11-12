package controllers

import (
	"net/http"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewExerciseType POST endpoint.
// Creates a new exercise types
func NewExerciseType(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	var newExerciseType types.ExerciseType
	if err := c.Bind(&newExerciseType); err != nil {
		log.Info(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newExerciseType.UserID = userClaim.ID

	if err := db.InsertNewExerciseType(newExerciseType); err != nil {
		log.Info(err)
		return c.JSON(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// GetExerciseTypes GET endpoint.
// Gets all exercise types belonging to the user that made the request
func GetExerciseTypes(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseTypes, err := db.GetExerciseTypesByUserID(userClaim.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exerciseTypes)
}

// EditExerciseTypes PUT endpoint
// @body types.ExerciseType
// Edit an existing record
// NOTE: ANY UNSENT DATA WILL WIPE THE FIELD ON THE
// NOTE: TODO:? Need to check for the fields that haven't been set and not update them to empty....maybe do that from the client side and always PUT all of the data unless the user wants to wipe it???...
func EditExerciseTypes(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	var userUpdate types.ExerciseType
	if err := c.Bind(&userUpdate); err != nil {
		log.Info(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userUpdate.UserID = userClaim.ID
	if userUpdate.ID.IsZero() {
		return c.String(http.StatusBadRequest, res.MissingID)
	}

	err := db.UpdateExerciseType(userClaim.ID, userUpdate)
	if err != nil {
		log.Info(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// DeleteExerciseType DELETE endpoint
// Delete an exercise type
// @params exercise_type_id
func DeleteExerciseType(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseTypesID := c.QueryParam("exercise_type_id")

	if exerciseTypesID == "" {
		return c.String(http.StatusBadRequest, res.MissingID)
	}

	exerciseTypesObjID, err := primitive.ObjectIDFromHex(exerciseTypesID)
	if err != nil {
		return c.String(http.StatusBadRequest, res.InternalServerError)
	}

	err = db.DeleteExerciseTypeByID(exerciseTypesObjID, userClaim.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, res.DatabseError)

	}
	return c.NoContent(http.StatusOK)
}
