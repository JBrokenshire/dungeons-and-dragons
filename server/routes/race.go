package routes

import (
	"dnd-api/db/stores"
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func raceRoutes(server *server.Server) {
	raceStore := stores.NewGormRaceStore(server.Db)
	raceController := controllers.RaceController{Store: raceStore}

	races := server.Echo.Group("/races")
	races.GET("", raceController.GetAll)
	races.GET("/:id", raceController.Get)
}
