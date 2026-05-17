package repository

import (
	"github.com/likhithkp/clip/data_access/repository/url"
	"github.com/likhithkp/clip/data_access/repository/user"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access-repository",
	fx.Provide(
		user.NewUserRepository,
		url.NewUrlRepository,
	),
)
