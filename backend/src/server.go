package src

import "github.com/labstack/echo"

func Server() {
	e := echo.New()
	e.Start(":8080")
}
