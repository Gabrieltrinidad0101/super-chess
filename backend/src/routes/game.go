package routes

import (
	"github.com/labstack/echo"
)

func Game(e *echo.Echo) {
	g := e.Group("/game")

	g.POST("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.GET("/:id", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})
}
