package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Phirimhel/go-todo-app/docs"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_midleware.Middleware
}

func NewHTTPserver(
	config Config,
	logger *core_logger.Logger,
	middleware ...core_http_midleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        logger,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterApiRoutes(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.ApiVersion)

		fmt.Println("version router path:", prefix)

		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleware()))
	}

}

func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	s.mux.HandleFunc(
		"/swagger/doc.json",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(docs.SwaggerInfo.ReadDoc()))

		},
	)
}

func (s *HTTPServer) Run(ctx context.Context) error {

	mux := core_http_midleware.ChaneMiddleware(s.mux, s.middleware...)

	mux2 := http.NewServeMux()
	_ = mux2

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("Start HTTP serveer", zap.String("addr", s.config.Addr))
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("losten and server HTTP: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
		s.log.Warn("HTTP server stoped")
	}

	return nil
}
