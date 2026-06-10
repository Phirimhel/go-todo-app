package core_http_midleware

import (
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
)

func MockMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			log.Debug("<<< 🔵 handler in")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🔵 handler out")
		})
	}

}

func RouterMockMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			log.Debug("<<< 🟡 router in")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🟡 router out")
		})
	}
}

func RouterMockServerMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			log.Debug("<<< 🟣 server in")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🟣 server out")
		})
	}
}
