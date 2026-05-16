package application

import (
	"github.com/likhithkp/clip/application/auth"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	auth.Module,
)
