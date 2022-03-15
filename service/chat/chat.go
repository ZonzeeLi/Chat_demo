/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/13/2022
    @UpdateDate: 3/15/2022
    @Note:       Service层 ChatService服务管理
**/

package chat

import (
	"chat_demo/global"
	"chat_demo/model/chat"
	"chat_demo/model/response"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"github.com/gorilla/websocket"
)

// Manager 用户管理
var Manager = chat.ClientManager{
	Clients:    make(map[string]*chat.Client), // 连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *chat.Broadcast),    // 广播结构
	Register:   make(chan *chat.Client),       // 用户信息注册
	Reply:      make(chan *chat.Client),
	Unregister: make(chan *chat.Client), // 用户连接关闭
}

type ChatService struct{}

func (ChatService *ChatService) CHAT(chatReq chat.ChatReq_Ws, conn *websocket.Conn) {
	client := &chat.Client{
		ID:     chat.CreateId(chatReq.Uid, chatReq.ToUid),
		SendID: chat.CreateId(chatReq.ToUid, chatReq.Uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 用户注册到用户管理上
	Manager.Register <- client
	go Read(client)
	go Write(client)
}

func Read(client *chat.Client) {
	defer func() { // 避免忘记关闭，所以要加上close
		Manager.Unregister <- client
		_ = client.Socket.Close()
	}()
	for {
		client.Socket.PongHandler()
		sendMsg := new(chat.SendMsg)
		// _,msg,_:=client.Socket.ReadMessage()
		err := client.Socket.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错
		if err != nil {
			global.Chat_LOG.Error("[Read] sendMsg invalid, err:", zap.Error(err))
			Manager.Unregister <- client
			_ = client.Socket.Close()
			break
		}

		if sendMsg.Type == 1 {
			r1, _ := global.Chat_REDIS.Get(context.TODO(), client.ID).Result()
			r2, _ := global.Chat_REDIS.Get(context.TODO(), client.SendID).Result()
			if r1 >= "3" && r2 == "" { // 对方无应答状态下，单方发送消息限制为3条
				replyMsg := chat.ReplyMsg{
					Code:    int(response.CodeWSLimit),
					CodeMsg: response.CodeWSLimit.Msg(),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
				_, _ = global.Chat_REDIS.Expire(context.TODO(), client.ID, chat.UnConnectExpire).Result() // 防止重复骚扰，未建立连接刷新过期时间一个月
				continue
			} else {
				global.Chat_REDIS.Incr(context.TODO(), client.ID)
				_, _ = global.Chat_REDIS.Expire(context.TODO(), client.ID, chat.ConnectExpire).Result() // 防止过快“分手”，建立连接三个月过期
			}
			global.Chat_LOG.Info("[Read] sendMsg success ", zap.String("ClientID", client.ID), zap.String("Content", sendMsg.Content))
			Manager.Broadcast <- &chat.Broadcast{
				Client:  client,
				Message: []byte(sendMsg.Content),
			}
		} else if sendMsg.Type == 2 { //拉取历史消息
			timeT, err := strconv.Atoi(sendMsg.Content) // 传送来时间
			if err != nil {
				global.Chat_LOG.Error("[Read] Time parse failed", zap.Error(err))
				timeT = 999999999
			}
			results := FindHistory(global.Chat_CONFIG.MongoDb.Dbname, client.SendID, client.ID, int64(timeT), 10)
			if len(results) > 10 {
				results = results[:10]
			} else if len(results) == 0 {
				replyMsg := chat.ReplyMsg{
					Code:    int(response.CodeWSNoHistory),
					CodeMsg: response.CodeWSNoHistory.Msg(),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			for _, result := range results {
				replyMsg := chat.ReplyMsg{
					From:    result.From,
					Content: result.Msg,
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		} else if sendMsg.Type == 3 {
			results := FindUnreadMsg(global.Chat_CONFIG.MongoDb.Dbname, client.SendID, client.ID)
			if len(results) == 0 {
				replyMsg := chat.ReplyMsg{
					Code:    int(response.CodeWsNoUnread),
					CodeMsg: response.CodeWsNoUnread.Msg(),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			for _, result := range results {
				replyMsg := chat.ReplyMsg{
					From:    result.From,
					Content: result.Msg,
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func Write(client *chat.Client) {
	defer func() {
		_ = client.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				_ = client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			global.Chat_LOG.Info("[Write] receive msg", zap.String("ClientID", client.ID), zap.String("msg", string(message)))
			replyMsg := chat.ReplyMsg{
				Code:    int(response.CodeWSSuccessMsg),
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
