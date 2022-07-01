package exercise_type

import (
	"net/http"

	"Logger.Fitness/go-libs/common"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const path = "/exercise_types"

// DatabaseInterface ...
type DatabaseInterface interface {
	InsertNewExerciseType(exerciseType types.ExerciseType) error
	GetExerciseTypesByUserID(userID primitive.ObjectID) ([]types.ExerciseType, error)
	UpdateExerciseType(userID primitive.ObjectID, exerciseType types.ExerciseType) error
	DeleteExerciseTypeByName(exerciseTypeName string, userID primitive.ObjectID) error
}

// ExerciseTypeController ...
type ExerciseTypeController struct {
	database       DatabaseInterface
	authMiddleware echo.MiddlewareFunc
}

// NewExerciseTypeController ...
func NewExerciseTypeController(database DatabaseInterface, authMiddleware
echo.MiddlewareFunc) ExerciseTypeController {
	return ExerciseTypeController{
		database:       database,
		authMiddleware: authMiddleware,
	}
}

// Init ...
func (ctrl *ExerciseTypeController) Init(g *echo.Group) {
	group := g.Group(path)

	group.GET("", ctrl.GetExerciseTypes, ctrl.authMiddleware)
	group.POST("", ctrl.NewExerciseType, ctrl.authMiddleware)
	group.PUT("", ctrl.EditExerciseTypes, ctrl.authMiddleware)
	group.DELETE("", ctrl.DeleteExerciseType, ctrl.authMiddleware)
	group.POST("/sync", ctrl.SyncAll, ctrl.authMiddleware)

}

// NewExerciseType POST endpoint.
// Creates a new exercise types
func (ctrl *ExerciseTypeController) NewExerciseType(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	var newExerciseType types.ExerciseType
	if err := c.Bind(&newExerciseType); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newExerciseType.UserID = userClaim.ID

	if err := db.InsertNewExerciseType(newExerciseType); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// GetExerciseTypes GET endpoint.
// Gets all exercise types belonging to the user that made the request
func (ctrl *ExerciseTypeController) GetExerciseTypes(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseTypes, err := db.GetExerciseTypesByUserID(userClaim.ID)
	if err != nil {
		return err
	}

	// NOTE: returning empty array in case of no items
	if len(exerciseTypes) == 0 {
		return c.JSON(http.StatusOK, []types.ExerciseType{})
	}

	return c.JSON(http.StatusOK, exerciseTypes)
}

// EditExerciseTypes PUT endpoint
// @body types.ExerciseType
// Edit an existing record
func (ctrl *ExerciseTypeController) EditExerciseTypes(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	var userUpdate types.ExerciseType
	if err := c.Bind(&userUpdate); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userUpdate.UserID = userClaim.ID
	if userUpdate.ID.IsZero() {
		return c.String(http.StatusBadRequest, res.MissingID)
	}

	err := db.UpdateExerciseType(userClaim.ID, userUpdate)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// DeleteExerciseType DELETE endpoint
// Delete an exercise type
// @params exercise_type_id
// TODO: either disable deletion if type has been used or offer to delete the
// 	typedata also????
func (ctrl *ExerciseTypeController) DeleteExerciseType(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	exerciseTypeName := c.QueryParam("name")

	if exerciseTypeName == "" {
		return c.String(http.StatusBadRequest, res.MissingID)
	}

	err := db.DeleteExerciseTypeByName(exerciseTypeName, userClaim.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, res.DatabseError)

	}
	return c.NoContent(http.StatusOK)
}

// SyncAll POST endpoint
// Syncs users local collection with remote (backend)
// @body []ExerciseType
func (ctrl *ExerciseTypeController) SyncAll(c echo.Context) error {
	db := ctrl.database
	userClaim := c.Get("user").(*types.JwtClaim)

	log.Infof("ExerciseType SyncAll %s", userClaim.ID.Hex())

	var userData []types.ExerciseType
	if err := c.Bind(&userData); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	serverData, err := db.GetExerciseTypesByUserID(userClaim.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, res.DatabseError)
	}

	updates, toDelete := extrapolateChanges(userData, serverData)

	log.Infof("Updates: %+v", updates)
	log.Infof("To delete: %+v", toDelete)

	err = ctrl.insertUpdatesOrNew(updates, userClaim.ID)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, res.DatabseError)
	}

	err = ctrl.deleteToDelete(toDelete, userClaim.ID)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

func extrapolateChanges(userData, serverData []types.ExerciseType) ([]types.ExerciseType,
	[]string) {

	var updates []types.ExerciseType
	var toDelete []string

	for _, userItem := range userData {
		found := false
		for _, serverItem := range serverData {
			if serverItem.Name == userItem.Name {
				found = true
				if !common.CompareExerciseTypes(userItem, serverItem) {
					updates = append(updates, userItem)
				}
			}
		}

		if !found {
			toDelete = append(toDelete, userItem.Name)
		}
	}

	return updates, toDelete
}

func (ctrl *ExerciseTypeController) insertUpdatesOrNew(updates []types.ExerciseType,
	userID primitive.ObjectID) error {
	db := ctrl.database

	for _, u := range updates {
		if u.ID.Hex() == "" {
			err := db.InsertNewExerciseType(u)
			if err != nil {
				return err
			}
		} else {
			err := db.UpdateExerciseType(userID, u)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ctrl *ExerciseTypeController) deleteToDelete(toDelete []string,
	userID primitive.ObjectID) error {
	db := ctrl.database

	for _, d := range toDelete {
		err := db.DeleteExerciseTypeByName(d, userID)
		if err != nil {
			return err
		}
	}
	return nil
}
