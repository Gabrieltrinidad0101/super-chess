package routes

import (
	"backend/src/conf"
	"backend/src/structs"
	"backend/src/utils"

	"github.com/labstack/echo"
)

var eventBus = utils.GetEventBus()

func User(e *echo.Echo) {
	conf := conf.Init()
	playerService := conf.Player()
	g := e.Group("/user")

	g.POST("/login", func(c echo.Context) error {
		var player *structs.Player
		if err := c.Bind(player); err != nil {
			return c.JSON(400, map[string]string{"message": "Invalid request"})
		}
		response := playerService.Login(player)
		return c.JSON(200, response)
	})

	g.POST("/register", func(c echo.Context) error {
		var player *structs.Player
		if err := c.Bind(player); err != nil {
			return c.JSON(400, map[string]string{"message": "Invalid request"})
		}
		response := playerService.Register(player)
		return c.JSON(200, response)
	})

	g.GET("/:id", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	g.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})
}
