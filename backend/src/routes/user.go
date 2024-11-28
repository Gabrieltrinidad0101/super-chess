package routes

import (
	"github.com/labstack/echo"
)

func User(e *echo.Echo) {
	g := e.Group("/user")

	g.POST("/login", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.POST("/register", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.GET("/:id", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})
}
