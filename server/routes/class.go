package routes

import (
	"dnd-api/db/stores"
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func classRoutes(server *server.Server) {
	classStore := stores.NewGormClassStore(server.Db)
	classController := controllers.ClassController{Store: classStore}

	classes := server.Echo.Group("/classes")
	classes.GET("", classController.GetAll)
	classes.GET("/:id", classController.Get)
	classes.PUT("/:id", classController.Update)
}
