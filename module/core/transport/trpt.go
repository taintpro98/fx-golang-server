package transport

import (
	"fmt"
	"fx-golang-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz business.IAuthenticateBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
) *Transport {
	fmt.Print("transport", authBiz)
	return &Transport{
		authBiz: authBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}

func (t *Transport) Register(ctx *gin.Context) {

}

func (t *Transport) Login(ctx *gin.Context) {

}

func (t *Transport) ListMovies(ctx *gin.Context) {

}
