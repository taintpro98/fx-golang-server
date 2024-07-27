package route

import (
	"fx-golang-server/module/core/transport"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	engine *gin.Engine,
	trpt *transport.Transport,
) {
	v1Api := engine.Group("/v1")

	publicApi := v1Api.Group("/public")
	publicApi.POST("/register", trpt.Register)
}
