/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/11/2022
    @UpdateDate: 3/15/2022
    @Note:       Service层 ServiceGroup封装暴露给Api层
**/

package service

import (
	"chat_demo/service/chat"
	"chat_demo/service/register"
)

type ServiceGroup struct {
	// You can also use Group here for further encapsulation
	RegisterService register.RegisterService
	ChatService     chat.ChatService
}

var ServiceGroupApp = new(ServiceGroup)
