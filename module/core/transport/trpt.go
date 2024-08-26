package transport

import (
	"fx-golang-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz  business.IAuthenticateBiz
	movieBiz business.IMovieBiz
	customerBiz business.ICustomerBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
	movieBiz business.IMovieBiz,
	customerBiz business.ICustomerBiz,
) *Transport {
	return &Transport{
		authBiz:  authBiz,
		movieBiz: movieBiz,
		customerBiz: customerBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}