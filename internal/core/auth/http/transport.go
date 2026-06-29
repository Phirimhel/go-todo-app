package core_auth_http

type authHTTPHandler struct {
	authService AuthService
}

func NewauthHTTPHandler(servise AuthService) *authHTTPHandler {
	return &authHTTPHandler{
		authService: servise,
	}
}
