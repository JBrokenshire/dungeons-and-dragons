package routes

import (
	"dungeons-and-dragons/db/stores"
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/server/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	classStore := stores.NewGormClassStore(server.Db)
	raceStore := stores.NewGormRaceStore(server.Db)
	classController := controllers.ClassController{Store: classStore}
	raceController := controllers.RaceController{Store: raceStore}
	characterController := controllers.CharacterController{
		CharacterStore: stores.NewGormCharacterStore(server.Db),
		ClassStore:     classStore,
		RaceStore:      raceStore,
	}

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
