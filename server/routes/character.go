package routes

import (
	"dnd-api/db/stores"
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func charactersRoutes(server *server.Server) {
	classStore := stores.NewGormClassStore(server.Db)
	raceStore := stores.NewGormRaceStore(server.Db)
	characterController := controllers.CharacterController{
		CharacterStore: stores.NewGormCharacterStore(server.Db),
		ClassStore:     classStore,
		RaceStore:      raceStore,
	}

	characters := server.Echo.Group("/characters")
	characters.GET("", characterController.GetAll)
	characters.POST("", characterController.Create)

	characters.GET("/:id", characterController.Get)
	characters.PUT("/:id", characterController.Update)
	characters.DELETE("/:id", characterController.Delete)
	characters.GET("/:id/inspiration", characterController.ToggleInspiration)
	characters.GET("/:id/level-up", characterController.LevelUp)
	characters.GET("/:id/heal/:value", characterController.Heal)
	characters.GET("/:id/damage/:value", characterController.Damage)
}
