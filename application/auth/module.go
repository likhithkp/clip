package auth

import (
	"github.com/likhithkp/clip/application/auth/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module("application-auth",
	fx.Provide(
		handlers.NewSignUpHandler,
		handlers.NewSignInHandler,
		NewAuthController,
	),
	fx.Invoke(RegisterAuthController),
)
