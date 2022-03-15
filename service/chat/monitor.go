/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/14/2022
    @UpdateDate: 3/15/2022
    @Note:       WebSocket服务监听
**/

package chat

import (
	"chat_demo/global"
	"chat_demo/model/chat"
	"chat_demo/model/response"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
)

func Monitor() {
	for {
		select {
		case conn := <-Manager.Register: // 建立连接
			global.Chat_LOG.Info("[Monitor] new connection :", zap.String("ConnectID", conn.ID))
			Manager.Clients[conn.ID] = conn
			replyMsg := &chat.ReplyMsg{
				Code:    int(response.CodeWSSuccess),
				CodeMsg: response.CodeWSSuccess.Msg(),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister: // 断开连接
			global.Chat_LOG.Info("[Monitor] bad connection :", zap.String("ConnectID", conn.ID))
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &chat.ReplyMsg{
					Code:    int(response.CodeWSEnd),
					CodeMsg: response.CodeWSEnd.Msg(),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		//广播信息
		case broadcast := <-Manager.Broadcast: // 1->2 进行广播
			message := broadcast.Message
			sendId := broadcast.Client.SendID       // 2
			flag := false                           // 默认对方不在线
			for id, conn := range Manager.Clients { // 去用户管理里找2是否在线
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID // 1
			if flag {
				global.Chat_LOG.Info("[Monitor] Online Reply")
				replyMsg := &chat.ReplyMsg{
					Code:    int(response.CodeWSOnlineReply),
					CodeMsg: response.CodeWSOnlineReply.Msg(),
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(global.Chat_CONFIG.MongoDb.Dbname, id, string(message), 1, int64(3*chat.Month))
				if err != nil {
					global.Chat_LOG.Error("[Monitor] InsertOneMsg failed ", zap.Error(err))
				}
			} else {
				global.Chat_LOG.Info("[Monitor] Offline Reply")
				replyMsg := chat.ReplyMsg{
					Code:    int(response.CodeWSOfflineReply),
					CodeMsg: response.CodeWSOfflineReply.Msg(),
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(global.Chat_CONFIG.MongoDb.Dbname, id, string(message), 0, int64(3*chat.Month))
				if err != nil {
					global.Chat_LOG.Error("[Monitor] InsertOneMsg failed ", zap.Error(err))
				}
			}
		}
	}
}
