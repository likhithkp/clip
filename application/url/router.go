package url

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUrlRoutes(app *fiber.App, controller *Controller, middleware fiber.Handler) {
	appGroup := app.Group("api/v1/urls")

	appGroup.Use(middleware)
	appGroup.Post("", controller.CreateUrl)
	appGroup.Get("/:code", controller.GetUrl)
}
