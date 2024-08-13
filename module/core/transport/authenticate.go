package transport

import (
	"fx-golang-server/module/core/dto"

	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Register a new user and return a token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   dto.CreateUserRequest  body  string  true  "CreateUserRequest"
// @Success 200 {string} string "token"
// @Failure 400 {object} map[string]interface{}
// @Router /v1/public/register [post]
func (t *Transport) Register(ctx *gin.Context) {
	var data dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	result, err := t.authBiz.Register(ctx, data)
	dto.HandleResponse(ctx, result, err)
}

// @Summary Login a new user
// @Description Login a new session
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   dto.CreateUserRequest  body  string  true  "CreateUserRequest"
// @Success 200 {string} string "token"
// @Failure 400 {object} map[string]interface{}
// @Router /v1/public/login [post]
func (t *Transport) Login(ctx *gin.Context) {
	var data dto.LoginRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	result, err := t.authBiz.Login(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
