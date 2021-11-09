package controllers

import (
	"net/http"

	"Logger.Fitness/backend/db"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// NewExerciseType creates a new exerciseTypes
func NewExerciseType(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	var newExerciseType types.ExerciseType
	if err := c.Bind(&newExerciseType); err != nil {
		log.Info(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newExerciseType.UserId = userClaim.ID

	if err := db.InsertNewExerciseType(newExerciseType); err != nil {
		log.Info(err)
		return c.JSON(http.StatusInternalServerError, res.DATABASE_ERR)
	}

	return c.NoContent(http.StatusOK)
}

func GetExerciseTypes(c echo.Context) error {
	db := c.Get("db").(*db.DbClient)
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseTypes, err := db.GetExerciseTypesByUserID(userClaim.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, exerciseTypes)
}

func EditExerciseTypes(c echo.Context) error {
	return nil
}

func DeleteExerciseType(c echo.Context) error {
	return nil
}
