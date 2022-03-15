/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/13/2022
    @UpdateDate: 3/15/2022
    @Note:       Api层 ApiGroup管理，将所有api封装给路由层
**/

package api

import (
	"chat_demo/api/chat"
	"chat_demo/api/register"
)

type ApiGroup struct {
	Register register.RegisterApi
	Chat     chat.ChatApi
}

var ApiGroupApp = new(ApiGroup)
