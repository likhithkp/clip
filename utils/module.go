package utils

import (
	"github.com/likhithkp/clip/utils/config"
	"github.com/likhithkp/clip/utils/jwt"
	"github.com/likhithkp/clip/utils/logger"
	"github.com/likhithkp/clip/utils/middleware"
	"github.com/likhithkp/clip/utils/other"
	"github.com/likhithkp/clip/utils/server"
	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	server.Module,
	logger.Module,
	config.Module,
	other.Module,
	jwt.Module,
	middleware.Module,
)
