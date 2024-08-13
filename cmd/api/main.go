package main

import (
	"context"
	"errors"
	"fmt"
	"fx-golang-server/config"
	"fx-golang-server/middleware"
	"fx-golang-server/module/core/business"
	"fx-golang-server/module/core/repository"
	"fx-golang-server/module/core/transport"
	"fx-golang-server/pkg/cache"
	"fx-golang-server/pkg/database"
	"fx-golang-server/pkg/telegram"
	"fx-golang-server/pkg/tracing"
	"fx-golang-server/route"
	"fx-golang-server/token"
	"net/http"
	"time"

	_ "fx-golang-server/docs" // Import generated docs

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func handleConnection(lc fx.Lifecycle, redisClient cache.IRedisClient) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info().Ctx(ctx).Msg("Closing connection...")
			if redisClient != nil {
				redisClient.CloseConnection()
			}
			return nil
		},
	})
}

var ConnectionModule = fx.Module(
	"connection",
	fx.Provide(
		database.PostgresqlDatabaseProvider,
		cache.RedisClientProvider,
		telegram.TelegramBotProvider,
	),
	fx.Invoke(handleConnection),
)

func NewGinEngine(trpt *transport.Transport, jwtMaker token.IJWTMaker) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
	)

	// Register the Swagger handler
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.RegisterHealthCheckRoute(engine)
	route.RegisterRoutes(engine, trpt, jwtMaker)
	return engine
}

func startHttp(lc fx.Lifecycle, cnf *config.Config, engine *gin.Engine, telegramBot telegram.ITelegramBot) {
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if telegramBot != nil {
				go telegramBot.GetMessages(ctx)
			}

			go func() {
				log.Info().Ctx(ctx).Msg(fmt.Sprintf("Running API on port %s...", cnf.AppInfo.ApiPort))
				tracing.InitLogger("api-service")

				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Ctx(ctx).Err(err).Msg("Run app error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Shutting down server...")
			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Error().Ctx(timeoutCtx).Err(err).Msg("Error shutting down server")
			} else {
				log.Info().Ctx(timeoutCtx).Msg("Server shutdown complete.")
			}
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
		fx.Provide(token.NewJWTMaker),
		repository.RepositoryModule,
		business.BusinessModule,
		fx.Provide(transport.NewTransport),
		fx.Provide(NewGinEngine),
		fx.Invoke(startHttp),
	)
	appFx.Run()
}
