package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_transport_http "github.com/Phirimhel/go-todo-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	usersRouters := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersRouters...)

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("faled to init application logger: %w", err)
		os.Exit(1)
	}
	defer logger.CloseFile()

	logger.Debug("starting TODO application")

	httpServer := core_http_server.NewHTTPserver(
		core_http_server.NewConfigMust(),
		logger,
		core_http_midleware.RequestID(),
		core_http_midleware.Logger(logger),
		core_http_midleware.Panic(),
		core_http_midleware.Trace(),
	)

	httpServer.RegisterApiRoutes(apiVersionRouter)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
