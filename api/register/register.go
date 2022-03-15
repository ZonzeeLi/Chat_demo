/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/11/2022
    @UpdateDate: 3/13/2022
    @Note:       Api层RegisterApi管理
**/

package register

import (
	"chat_demo/global"
	"chat_demo/model/register"
	"chat_demo/model/response"
	"chat_demo/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterApi struct{}

var registerService = service.ServiceGroupApp.RegisterService

func (registerApi *RegisterApi) Register(c *gin.Context) {
	var userRegister register.UserRegisterRequest
	if err := c.ShouldBind(&userRegister); err != nil {
		global.Chat_LOG.Error("Invalid Params, err:", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParams)
	} else {
		res, msg := registerService.Register(userRegister)
		response.ResponseSet(c, res, msg)
	}
}
