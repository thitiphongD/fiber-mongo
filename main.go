package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/fiber-mongo/configs"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Listen(":6000")
}
