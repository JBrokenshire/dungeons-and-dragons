package routes

import (
	"dungeons-and-dragons/db/stores"
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/server/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	classController := controllers.ClassController{Store: stores.NewGormClassStore(server.Db)}
	raceController := controllers.RaceController{Store: stores.NewGormRaceStore(server.Db)}
	characterController := controllers.CharacterController{Store: stores.NewGormCharacterStore(server.Db)}

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
