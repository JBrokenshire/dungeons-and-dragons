package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func itemRoutes(server *server.Server) {
	itemController := controllers.ItemController{
		Store: server.Stores.ItemStore,
	}

	items := server.Echo.Group("/items")
	items.GET("", itemController.GetAll)
}

func weaponRoutes(server *server.Server) {
	weaponController := controllers.WeaponController{
		Store: server.Stores.WeaponStore,
	}

	weapons := server.Echo.Group("/items/weapons")
	weapons.GET("", weaponController.GetAll)
}
