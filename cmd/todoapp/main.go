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
	web_fs_repository "github.com/Phirimhel/go-todo-app/internal/features/web/repository/file_system"
	web_service "github.com/Phirimhel/go-todo-app/internal/features/web/service"
	web_transport_http "github.com/Phirimhel/go-todo-app/internal/features/web/trasport/http"
	"go.uber.org/zap"

	_ "github.com/Phirimhel/go-todo-app/docs"
)

// @title 			Golang Todo API
// @version 		1.0
// @description 	This is a production-ready RESTful Todo API server written in Go.
// @host 			127.0.0.1:8080
// @BasePath 		/api/v1
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

	//timezone
	logger.Debug("application time zone", zap.String("zone:", time.Local.String()))

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

	//main page html
	logger.Debug("initializing features", zap.String("feature", "web"))
	webRepository := web_fs_repository.NewWebRepository()
	webService := web_service.NewWebService(webRepository)
	webHTTPHandler := web_transport_http.NewWebHTTPHandler(webService)

	// server config
	logger.Debug("initializing HTTP server")
	serverConfig := core_http_server.NewConfigMust()
	logger.Warn("allowed origins:", zap.Any("origins", serverConfig.AlllowedOrigins))
	httpServer := core_http_server.NewHTTPserver(
		serverConfig,
		logger,
		// server middleware
		core_http_midleware.CORS(serverConfig.AlllowedOrigins),
		core_http_midleware.RequestID(),
		core_http_midleware.Logger(logger),
		core_http_midleware.Trace(),
		core_http_midleware.Panic(),
	)

	// routers (V1)
	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(statsTransportHTTP.Routes()...)
	// server routes
	httpServer.RegisterRoutes(webHTTPHandler.Routes()...)
	httpServer.RegisterApiRoutes(apiVersionRouter)
	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
