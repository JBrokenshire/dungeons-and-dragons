package main

import (
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/server/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := server.NewServer()

	routes.ConfigureRoutes(app)

	err = app.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Port already in use")
	}
}
