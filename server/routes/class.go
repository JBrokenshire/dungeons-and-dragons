package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func classRoutes(server *server.Server) {
	classController := controllers.ClassController{Store: server.Stores.Class}

	classes := server.Echo.Group("/classes")
	classes.GET("", classController.GetAll)
	classes.GET("/:id", classController.Get)
	classes.PUT("/:id", classController.Update)
}
