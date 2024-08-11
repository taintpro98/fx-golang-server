package transport

import (
	"fx-golang-server/module/core/dto"

	"github.com/gin-gonic/gin"
)

func (t *Transport) Register(ctx *gin.Context) {
	var data dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	result, err := t.authBiz.Register(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
