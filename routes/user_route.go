package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/fiber-mongo/controllers"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}
