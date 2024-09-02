package routes

import (
	"dungeons-and-dragons/db/stores"
	"dungeons-and-dragons/server"
	controller2 "dungeons-and-dragons/server/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	classController := controller2.ClassController{Store: stores.NewGormClassStore(server.Db)}
	raceController := controller2.RaceController{Store: stores.NewGormRaceStore(server.Db)}
	characterController := controller2.CharacterController{Store: stores.NewGormCharacterStore(server.Db)}

	server.Echo.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "hello", "Jared")
	})

	server.Echo.GET("/classes", classController.GetAll)
	server.Echo.GET("/classes/:id", classController.Get)
	server.Echo.PUT("/classes/:id", classController.Update)

	server.Echo.GET("/races", raceController.GetAll)
	server.Echo.GET("/races/:id", raceController.Get)

	server.Echo.GET("/characters", characterController.GetAll)
	server.Echo.POST("/characters", characterController.Create)

	server.Echo.GET("/characters/:id", characterController.Get)
	server.Echo.PUT("/characters/:id", characterController.Update)
	server.Echo.DELETE("/characters/:id", characterController.Delete)

	server.Echo.PUT("/characters/:id/level-up", characterController.LevelUp)
}
