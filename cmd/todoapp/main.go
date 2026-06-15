package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Phirimhel/go-todo-app/internal/core/config"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_pgx_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool/pgx"
	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	statistics_postgres_repository "github.com/Phirimhel/go-todo-app/internal/features/statistics/repository/postgres"
	statistic_service "github.com/Phirimhel/go-todo-app/internal/features/statistics/service"
	statistics_transport_http "github.com/Phirimhel/go-todo-app/internal/features/statistics/transport/http"
	tasks_repository "github.com/Phirimhel/go-todo-app/internal/features/tasks/repository/postgres"
	tasks_service "github.com/Phirimhel/go-todo-app/internal/features/tasks/service"
	tasks_transport_http "github.com/Phirimhel/go-todo-app/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/Phirimhel/go-todo-app/internal/features/users/repository/postgres"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
	users_transport_http "github.com/Phirimhel/go-todo-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	// time zone
	globalConfig := core_config.NewGlobalConfigMust()
	time.Local = globalConfig.TimeZone

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

	//timezone
	logger.Debug("application time zone", zap.String("zone:", time.Local.String()))

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

	//tasks
	logger.Debug("initializing features", zap.String("feature", "tasks"))
	tasksRepository := tasks_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	//statistics
	logger.Debug("initializing features", zap.String("feature", "statistics"))
	statsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statsService := statistic_service.NewStatisticsSercice(statsRepository)
	statsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statsService)

	// server config
	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPserver(
		core_http_server.NewConfigMust(),
		logger,
		// server middleware
		core_http_midleware.RequestID(),
		core_http_midleware.Logger(logger),
		core_http_midleware.Trace(),
		core_http_midleware.RouterMockServerMiddleware(),
		core_http_midleware.Panic(),
	)

	// routers (V1)
	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(statsTransportHTTP.Routes()...)

	// routers (V2 with middlerware)
	apiVersionRouterV2 := core_http_server.NewApiVersionRouter(
		core_http_server.ApiVersion2,
		core_http_midleware.RouterMockMiddleware(),
	)
	apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterApiRoutes(apiVersionRouter, apiVersionRouterV2)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
