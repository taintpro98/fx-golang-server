package route

import (
	trpt "fx-golang-server/module/core/transport"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckRoute(e *gin.Engine) {
	e.GET("/health", trpt.HandleHealthCheck)
}
