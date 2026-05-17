package auth

import "github.com/gofiber/fiber/v2"

func RegisterAuthController(app *fiber.App, controller *Controller) {
	appGroup := app.Group("api/v1/auth")

	appGroup.Post("sign-up", controller.SignUpHandler)
	appGroup.Post("login", controller.SignInHandler)
}
