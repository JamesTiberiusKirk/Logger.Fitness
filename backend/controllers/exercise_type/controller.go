package exercise_type

import (
	"net/http"

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
	DeleteExerciseTypeByID(exerciseTypeID, userID primitive.ObjectID) error
}

// ExerciseTypeController ...
type ExerciseTypeController struct {
	database       DatabaseInterface
	authMiddleware echo.MiddlewareFunc
}

// NewExerciseTypeController ...
func NewExerciseTypeController(database DatabaseInterface, authMiddleware echo.MiddlewareFunc) ExerciseTypeController {
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

	// TODO: Find a way to omit user_id from the response json
	//for i := 0; i < len(exerciseTypes); i++ {
	//exerciseTypes[i].UserID = primitive.NilObjectID
	//}

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

	exerciseTypesID := c.QueryParam("id")

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
