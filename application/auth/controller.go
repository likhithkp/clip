package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/auth/handlers"
)

type AuthController struct {
	signUpHandler *handlers.SignUpHandler
	signInHandler *handlers.SignInHandler
}

func NewAuthController(
	signUpHandler *handlers.SignUpHandler,
	signInHandler *handlers.SignInHandler,
) *AuthController {
	return &AuthController{
		signUpHandler: signUpHandler,
		signInHandler: signInHandler,
	}
}

func (controller *AuthController) SignUpHandler(ctx *fiber.Ctx) error {
	return controller.signUpHandler.SignUp(ctx)
}

func (controller *AuthController) SignInHandler(ctx *fiber.Ctx) error {
	return controller.signInHandler.SignIn(ctx)
}
