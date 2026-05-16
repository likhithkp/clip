package other

import (
	"go.uber.org/fx"
)

var Module = fx.Module("utils-others",
	fx.Provide(NewResponseStruct),
)
