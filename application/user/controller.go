package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/application/user/handlers"
)

type UserController struct {
	getUserDetails *handlers.GetUserDetailsHandler
}

func NewUserController(
	getUserDetails *handlers.GetUserDetailsHandler,
) *UserController {
	return &UserController{
		getUserDetails: getUserDetails,
	}
}

func (controller *UserController) GetUserDetails(ctx *fiber.Ctx) error {
	return controller.getUserDetails.GetUserDetails(ctx)
}
