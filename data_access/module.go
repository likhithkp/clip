package data_access

import (
	"github.com/likhithkp/clip/data_access/mongo"
	redisClient "github.com/likhithkp/clip/data_access/redis"
	"github.com/likhithkp/clip/data_access/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access",
	mongo.Module,
	repository.Module,
	redisClient.Module,
)
