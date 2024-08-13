package transport

import (
	"fx-golang-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz  business.IAuthenticateBiz
	movieBiz business.IMovieBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
	movieBiz business.IMovieBiz,
) *Transport {
	return &Transport{
		authBiz:  authBiz,
		movieBiz: movieBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}