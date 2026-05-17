package url

import (
	"github.com/likhithkp/clip/application/url/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module("application-url",
	fx.Provide(
		handlers.NewCreateUrlHanler,
		handlers.NewGetUrlHandler,
		NewController,
	),
	fx.Invoke(RegisterUrlRoutes),
)
