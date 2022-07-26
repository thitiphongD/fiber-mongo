package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/fiber-mongo/configs"
	"github.com/thitiphongD/fiber-mongo/routes"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
