package business

import "go.uber.org/fx"

var BusinessModule = fx.Module(
	"business",
	fx.Provide(NewAuthenticateBiz, NewMovieBiz),
)
