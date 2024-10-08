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
	proficienciesController := controllers.CharacterProficienciesController{CharacterProficienciesStore: server.Stores.CharacterProficiencies}
	defensesController := controllers.CharacterDefensesController{CharacterDefensesStore: server.Stores.CharacterDefensesStore}
	conditionsController := controllers.CharacterConditionsController{CharacterConditionsStore: server.Stores.CharacterConditionsStore}
	inventoryController := controllers.CharacterInventoryController{Store: server.Stores.CharacterInventoryStore}
	moneyController := controllers.CharacterMoneyController{Store: server.Stores.CharacterMoneyStore}

	characters := server.Echo.Group("/characters")
	characters.GET("", characterController.GetAll)
	characters.POST("", characterController.Create)

	characters.GET("/:id", characterController.Get)
	characters.PUT("/:id", characterController.Update)
	characters.DELETE("/:id", characterController.Delete)
	characters.GET("/:id/inspiration", characterController.ToggleInspiration)
	characters.GET("/:id/level-up", characterController.LevelUp)
	characters.PUT("/:id/heal/:value", characterController.Heal)
	characters.PUT("/:id/damage/:value", characterController.Damage)

	characters.GET("/:id/proficient-skills", skillsController.GetProficientByCharacterID)
	characters.GET("/:id/senses", sensesController.GetSensesByCharacterID)

	characters.GET("/:id/proficient/armour", proficienciesController.GetCharacterProficientArmourTypes)
	characters.GET("/:id/proficient/weapons", proficienciesController.GetCharacterProficientWeapons)
	characters.GET("/:id/proficient/tools", proficienciesController.GetCharacterProficientTools)
	characters.GET("/:id/proficient/languages", proficienciesController.GetCharacterLanguages)

	characters.GET("/:id/defenses", defensesController.GetCharacterDefenses)
	characters.GET("/:id/conditions", conditionsController.GetCharacterConditions)

	characters.GET("/:id/inventory", inventoryController.GetCharacterInventory)
	characters.GET("/:id/inventory/equipped-weapons", inventoryController.GetCharacterEquippedWeapons)
	characters.GET("/:id/inventory/money", moneyController.GetCharacterMoney)
	characters.PUT("/:characterID/inventory/:itemID", inventoryController.ToggleItemEquipped)
}
