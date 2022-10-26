package workouttypes

import (
	"Logger.Fitness/go-libs/types"
	"github.com/labstack/echo"
)

const path = "/workout_templates"

// DatabaseInterface ...
type DatabaseInterface interface {
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

}

// DTO - workout & their exercises Data Transfer Object
type DTO struct {
	Workout   types.Workout    `json:"workout"`
	Exercises []types.Exercise `json:"exercises"`
}
