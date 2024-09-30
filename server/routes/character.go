package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func charactersRoutes(server *server.Server) {
	characterController := controllers.CharacterController{
		CharacterStore: server.Stores.Character,
		ClassStore:     server.Stores.Class,
		RaceStore:      server.Stores.Race,
	}
	skillsController := controllers.CharacterSkillsController{CharacterSkillsStore: server.Stores.CharacterSkills}
	sensesController := controllers.CharacterSensesController{CharacterSensesStore: server.Stores.CharacterSenses}
	proficientArmourController := controllers.CharacterArmourTypesController{CharacterArmourTypesStore: server.Stores.CharacterArmourTypes}

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

	characters.GET("/:id/proficient-skills", skillsController.GetProficientByCharacterID)
	characters.GET("/:id/senses", sensesController.GetSensesByCharacterID)
	characters.GET("/:id/proficient-armour", proficientArmourController.GetProficientArmourTypes)
}
