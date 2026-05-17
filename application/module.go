package application

import (
	"github.com/likhithkp/clip/application/auth"
	"github.com/likhithkp/clip/application/url"
	"github.com/likhithkp/clip/application/user"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	auth.Module,
	url.Module,
	user.Module,
)
