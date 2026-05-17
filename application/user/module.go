package user

import (
	"github.com/likhithkp/clip/application/user/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module("application-user",
	fx.Provide(
		handlers.NewGetUserDetailsHandler,
		NewUserController,
	),
	fx.Invoke(RegisterUserRoutes),
)
