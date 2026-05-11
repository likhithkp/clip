package main

import (
	"github.com/likhithkp/clip/application"
	"github.com/likhithkp/clip/data_access"
	"github.com/likhithkp/clip/utils"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		application.Module,
		data_access.Module,
		utils.Module,
	).Run()
}
