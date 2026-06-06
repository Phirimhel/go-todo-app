package main

import (
	"context"
	"os/signal"
	"syscall"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_pgx_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool/pgx"
	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_postgres_repository "github.com/Phirimhel/go-todo-app/internal/features/users/repository/postgres"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
	users_transport_http "github.com/Phirimhel/go-todo-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	// main ctx
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	// logger
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		logger.Fatal("faled to init application logger:", zap.Error(err))
	}
	defer logger.CloseFile()

	// conn pool
	logger.Debug("initializing postgres conection pool")
	pool, err := core_pgx_pool.NewPgxConnectionPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres conection pool:", zap.Error(err))
	}
	defer pool.Close()

	// users
	logger.Debug("initializing features", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewRepository(pool)
	userService := users_service.NewUserService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(userService)

	// server config
	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPserver(
		core_http_server.NewConfigMust(),
		logger,
		core_http_midleware.RequestID(),
		core_http_midleware.Logger(logger),
		core_http_midleware.Trace(),
		core_http_midleware.Panic(),
	)

	// routers
	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterApiRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
