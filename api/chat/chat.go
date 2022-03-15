/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/13/2022
    @UpdateDate: 3/15/2022
    @Note:       Api层 ChatApi管理
**/

package chat

import (
	"chat_demo/global"
	"chat_demo/model/chat"
	"chat_demo/service"
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

type ChatApi struct{}

var chatService = service.ServiceGroupApp.ChatService

func (chatApi *ChatApi) CHAT(c *gin.Context) {
	chatReq := chat.ChatReq_Ws{
		Uid:   c.Query("uid"),
		ToUid: c.Query("touid"),
	}

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		global.Chat_LOG.Error("[WebSocket] Upgrade failed", zap.Error(err))
		http.NotFound(c.Writer, c.Request)
		return
	}
	chatService.CHAT(chatReq, conn)

}
