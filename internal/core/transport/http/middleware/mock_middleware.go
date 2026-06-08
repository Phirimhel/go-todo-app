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

			log.Debug("<<< 🔵DUMMY LOG IN")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🟡DUMMY LOG OUT")
		})
	}

}

func RouterMockMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			log.Debug("<<< 🔵DUMMY ROUTER LOG IN")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🟡DUMMY ROUTER LOG OUT")
		})
	}
}

func RouterMockServerMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			log.Debug("<<< 🟣DUMMY SERVER OBAL LOG IN")
			next.ServeHTTP(w, r)
			log.Debug(">>> 🟣DUMMY SERVER LOG OUT")
		})
	}
}
