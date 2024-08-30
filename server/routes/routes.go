package routes

import (
	"dungeons-and-dragons/controller"
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	server.Echo.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "hello", "Jared")
	})

	classController := controller.ClassController{Store: store.NewGormClassStore()}
	server.Echo.GET("/classes", classController.GetAll)
	server.Echo.GET("/classes/:id", classController.Get)
	server.Echo.PUT("/classes/:id", classController.Update)

	raceController := controller.RaceController{Store: store.NewGormRaceStore()}
	server.Echo.GET("/races", raceController.GetAll)
	server.Echo.GET("/races/:id", raceController.Get)

	characterController := controller.CharacterController{Store: store.NewGormCharacterStore()}
	server.Echo.GET("/characters", characterController.GetAll)
	server.Echo.POST("/characters", characterController.Create)

	server.Echo.GET("/characters/:id", characterController.Get)
	server.Echo.PUT("/characters/:id", characterController.Update)
	server.Echo.DELETE("/characters/:id", characterController.Delete)

	server.Echo.PUT("/characters/:id/level-up", characterController.LevelUp)
}
