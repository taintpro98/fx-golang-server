package route

import (
	"fx-golang-server/middleware"
	"fx-golang-server/module/core/transport"
	"fx-golang-server/token"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	engine *gin.Engine,
	trpt *transport.Transport,
	jwtMaker token.IJWTMaker,
) {
	v1Api := engine.Group("/v1")

	publicApi := v1Api.Group("/public")
	publicApi.POST("/register", trpt.Register)
	publicApi.POST("/login", trpt.Login)

	publicApi.Use(middleware.AuthMiddleware(jwtMaker))
	{
		movieApi := publicApi.Group("/movies")
		{
			movieApi.GET("", trpt.ListMovies)
		}

		customerApi := publicApi.Group("/customer")
		{
			customerApi.GET("/profile", trpt.GetCustomerProfile)
		}
	}
}
