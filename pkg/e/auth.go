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
	ErrInvalidRefreshToken = CustomErr{
		HttpStatusCode: http.StatusUnauthorized,
		Code:           http.StatusUnauthorized,
		Msg:            "Refresh token is invalid",
		Language:       "refresh_token_is_invalid",
	}
)
