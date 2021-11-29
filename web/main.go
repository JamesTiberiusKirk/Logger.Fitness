package main

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "public")

	endless.DefaultHammerTime = 10 * time.Second
	endless.DefaultReadTimeOut = 295 * time.Second
	if err := endless.ListenAndServe(":8080", e); err != nil {
		log.Infof("Server stopped: %s", err)
	}
}
