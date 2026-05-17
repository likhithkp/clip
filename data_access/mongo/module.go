package mongo

import (
	"github.com/likhithkp/clip/data_access/mongo/url"
	"github.com/likhithkp/clip/data_access/mongo/user"
	"go.uber.org/fx"
)

var Module = fx.Module("mongo",
	fx.Provide(
		NewClient,
		NewDatabase,
		user.NewUserMongoService,
		url.NewUrlMongoService,
	),
)
