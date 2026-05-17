package redis

import (
	"github.com/likhithkp/clip/data_access/redis/url"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access-redis",
	fx.Provide(
		NewRedisClient,
		url.NewUrlRedisService,
	),
)
