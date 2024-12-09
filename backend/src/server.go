package src

import (
	"backend/src/routes"
	"backend/src/services"

	"github.com/labstack/echo"
)

func Server() {
	e := echo.New()
	services.Init()
	routes.User(e)
	e.Start(":8080")
}
