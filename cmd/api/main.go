package main

import (
	"context"
	"errors"
	"fmt"
	"fx-golang-server/config"
	"fx-golang-server/module/core/business"
	"fx-golang-server/module/core/repository"
	"fx-golang-server/module/core/transport"
	"fx-golang-server/pkg/cache"
	"fx-golang-server/pkg/database"
	"fx-golang-server/route"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var ConnectionModule = fx.Module(
	"connection",
	fx.Provide(
		database.PostgresqlDatabaseProvider,
		cache.RedisClientProvider,
	),
)

func NewGinEngine(trpt *transport.Transport) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	
	route.RegisterRoutes(engine, trpt)
	return engine
}

func startHttp(lc fx.Lifecycle, cnf *config.Config, engine *gin.Engine) {
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info().Ctx(ctx).Msg(fmt.Sprintf("Running API on port %s...", cnf.AppInfo.ApiPort))
				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Ctx(ctx).Err(err).Msg("Run app error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	appFx := fx.New(
		fx.Provide(
			func() *config.Config {
				cnf := config.Init()
				return &cnf
			},
		),
		ConnectionModule,
		repository.RepositoryModule,
		business.BusinessModule,
		fx.Provide(transport.NewTransport),
		fx.Provide(NewGinEngine),
		fx.Invoke(startHttp),
	)
	appFx.Run()
}
