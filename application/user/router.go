package user

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, controller *UserController, middleware fiber.Handler) {
	appGroup := app.Group("api/v1/users")

	appGroup.Use(middleware)
	appGroup.Get("", controller.GetUserDetails)
}
