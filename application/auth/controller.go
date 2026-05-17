package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/auth/handlers"
)

type Controller struct {
	signUpHandler *handlers.SignUpHandler
	signInHandler *handlers.SignInHandler
}

func NewController(
	signUpHandler *handlers.SignUpHandler,
	signInHandler *handlers.SignInHandler,
) *Controller {
	return &Controller{
		signUpHandler: signUpHandler,
		signInHandler: signInHandler,
	}
}

func (controller *Controller) SignUpHandler(ctx *fiber.Ctx) error {
	return controller.signUpHandler.SignUp(ctx)
}

func (controller *Controller) SignInHandler(ctx *fiber.Ctx) error {
	return controller.signInHandler.SignIn(ctx)
}
