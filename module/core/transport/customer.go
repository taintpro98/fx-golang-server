package transport

import (
	"fx-golang-server/module/core/dto"
	"fx-golang-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetCustomerProfile(ctx *gin.Context) {
	userID := ctx.MustGet(constants.XUserID).(string)
	profile, err := t.customerBiz.GetCustomerProfile(ctx, userID)
	dto.HandleResponse(ctx, profile, err)
}

func (t *Transport) SearchUsers(ctx *gin.Context) {
	var data dto.SearchUsersRequest
	if err := ctx.ShouldBindQuery(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	result, err := t.elasticBiz.SearchUsers(ctx, data)
	dto.HandleResponse(ctx, result, err)
}