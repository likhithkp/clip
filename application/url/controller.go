package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/url/handlers"
)

type Controller struct {
	createUrlHandler *handlers.CreateUrlHandler
	getUrlHandler    *handlers.GetUrlHandler
}

func NewController(
	createUrlHandler *handlers.CreateUrlHandler,
	getUrlHandler *handlers.GetUrlHandler,
) *Controller {
	return &Controller{
		createUrlHandler: createUrlHandler,
		getUrlHandler:    getUrlHandler,
	}
}

func (controller *Controller) CreateUrl(ctx *fiber.Ctx) error {
	return controller.createUrlHandler.CreateUrl(ctx)
}

func (controller *Controller) GetUrl(ctx *fiber.Ctx) error {
	return controller.getUrlHandler.GetUrl(ctx)
}
