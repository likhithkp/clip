package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/auth/handlers"
)

type AuthController struct {
	signUpHandler *handlers.SignUpHandler
}

func NewAuthController(
	signUpHandler *handlers.SignUpHandler,
) *AuthController {
	return &AuthController{
		signUpHandler: signUpHandler,
	}
}

func (controller *AuthController) SignUpHandler(ctx *fiber.Ctx) error {
	return controller.signUpHandler.SignUp(ctx)
}
