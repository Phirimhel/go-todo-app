package core_http_midleware

import (
	"context"
	"fmt"
	"net/http"

	core_auth "github.com/Phirimhel/go-todo-app/internal/core/auth"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

func JWT(authenticator core_auth.Authenticator) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := core_logger.MustGetLoggerFromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			fmt.Println("😊")

			Bearer, err := core_http_utils.GetBearerToken(r.Header)
			if err != nil {
				responseHandler.ErrorResponse(err, "Unauthorized")
				return
			}

			claims, err := authenticator.ValidateJWT(Bearer)
			if err != nil {
				responseHandler.ErrorResponse(err, "Unauthorized")
				return
			}

			ctx = context.WithValue(ctx, core_auth.ClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
