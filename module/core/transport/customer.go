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
