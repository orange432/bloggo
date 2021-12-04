package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/orange432/bloggo/pkg/handlers"
)

const PORT_NUMBER = ":3000"

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Load Routes
	app.Get("/", handlers.Home)
	app.Get("/about", handlers.About)
	app.Get("/articles/:slug", handlers.Article)

	app.Static("/", "./public")

	fmt.Println("Server started at http://localhost" + PORT_NUMBER)
	app.Listen(PORT_NUMBER)
}
