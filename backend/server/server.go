package server

import (
	"Logger.Fitness/backend/controllers"
	"Logger.Fitness/backend/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ContextParams struct {
	DbClient db.Client
	Port     string
}

func Run(params *ContextParams) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", controllers.HelloWorld)

	e.Logger.Fatal(e.Start(params.Port))
}
