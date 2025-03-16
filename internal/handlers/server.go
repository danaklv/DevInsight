package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer() *fiber.App {
	engine := html.New("../internal/front/templates/", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/styles", "../internal/front/styles/")

	return app
}
