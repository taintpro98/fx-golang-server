package transport

import (
	"fx-golang-server/module/core/business"
	"net/http"

	"fx-golang-server/module/core/dto"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz business.IAuthenticateBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
) *Transport {
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
	var data dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	result, err := t.authBiz.Register(ctx, data)
	dto.HandleResponse(ctx, result, err)
}

func (t *Transport) Login(ctx *gin.Context) {

}

func (t *Transport) ListMovies(ctx *gin.Context) {

}
