package core_http_midleware

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func CORS(allowedOriginsList []string) Middleware {

	fmt.Println("CORS")
	allowedOrigins := make(map[string]struct{})

	for _, origin := range allowedOriginsList {
		allowedOrigins[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				origin := r.Header.Get("Origin")
				path := r.URL.Path

				fmt.Println("origin", origin, "path", path)

				if _, ok := allowedOrigins[origin]; ok {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				}

				if r.Method == http.MethodOptions {
					w.WriteHeader(http.StatusOK)
					return
				}

				next.ServeHTTP(w, r)
			},
		)
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				requestID := r.Header.Get(requestIDHeader)

				if requestID == "" {
					requestID = uuid.NewString()
				}

				r.Header.Set(requestIDHeader, requestID)
				w.Header().Set(requestIDHeader, requestID)

				next.ServeHTTP(w, r)
			},
		)
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), core_logger.LogerContextKey, l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)

			before := time.Now().UTC()

			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", before),
			)

			wrapedResponseWriter := core_http_response.NewResponseWriter(w)

			next.ServeHTTP(wrapedResponseWriter, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("code", wrapedResponseWriter.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(before)),
			)

		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.GetLoggerFromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					stack := debug.Stack()
					responseHandler.PanicResponse(p, "durring handle HTTP request got unexpected panic", stack)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// func TestMiddlewareIN() Middleware {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			ctx := r.Context()
// 			log := core_logger.GetLoggerFromContext(ctx)

// 			log.Debug("test_nidle_ware_IN")
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// func TestMiddlewareOUT() Middleware {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			ctx := r.Context()
// 			log := core_logger.GetLoggerFromContext(ctx)

// 			log.Debug("test_nidle_ware_OUT")
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
