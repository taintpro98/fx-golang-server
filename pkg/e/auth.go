package e

import "net/http"

var (
	ErrUnauthorized = CustomErr{
		HttpStatusCode: http.StatusUnauthorized,
		Code:           http.StatusUnauthorized,
		Msg:            "Unauthorized",
		Language:       "unauthorized",
	}
	ErrTokenExpired = CustomErr{
		HttpStatusCode: http.StatusUnauthorized,
		Code:           http.StatusUnauthorized,
		Msg:            "Token is expired",
		Language:       "token_expired",
	}
)
