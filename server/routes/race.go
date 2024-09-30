package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func raceRoutes(server *server.Server) {
	raceController := controllers.RaceController{Store: server.Stores.Race}

	races := server.Echo.Group("/races")
	races.GET("", raceController.GetAll)
	races.GET("/:id", raceController.Get)
}
