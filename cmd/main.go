package main

import (
	"AiTaskManager/internal/config"
	"AiTaskManager/internal/handlers"
	"log"
)

func main() {
	app := handlers.StartServer()
	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

	config.InitDatabase()

	config.CreateTables()

}
