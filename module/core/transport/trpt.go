package transport

import (
	"fx-golang-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz     business.IAuthenticateBiz
	movieBiz    business.IMovieBiz
	customerBiz business.ICustomerBiz
	elasticBiz  business.IElasticBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
	movieBiz business.IMovieBiz,
	customerBiz business.ICustomerBiz,
	elasticBiz business.IElasticBiz,
) *Transport {
	return &Transport{
		authBiz:     authBiz,
		movieBiz:    movieBiz,
		customerBiz: customerBiz,
		elasticBiz:  elasticBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}
