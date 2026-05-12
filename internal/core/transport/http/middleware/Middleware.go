package core_http_midleware

import "net/http"

type Middleware func(http.Handler) http.Handler
