/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/13/2022
    @UpdateDate: 3/15/2022
    @Note:       Chat服务层结构
**/

package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	UnConnectExpire = time.Hour * 24 * 30
	ConnectExpire   = time.Hour * 24 * 30 * 3
	Month           = 60 * 60 * 24 * 30
)

// ChatReq_Ws 聊天请求结构
type ChatReq_Ws struct {
	Uid   string `json:"uid"`
	ToUid string `json:"toUid"`
}

// SendMsg 发送消息结构
type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReplyMsg 回复消息结构
type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	CodeMsg string `json:"code_msg"`
	//Content []byte `json:"content"`
	Content string `json:"content"`
}

// Client id用户连接结构
type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

// Broadcast 广播结构
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// ClientManager 用户管理
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

// Message 信息转JSON (包括：发送者、接收者、内容)
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

// Trainer MongoDb存储结构
type Trainer struct {
	Content   string `bson:"content"`   // 内容
	StartTime int64  `bson:"startTime"` // 创建时间
	EndTime   int64  `bson:"endTime"`   // 过期时间
	Read      uint   `bson:"read"`      // 已读
}

// Result 响应结构
type Result struct {
	StartTime int64
	Msg       string
	//Msg     []byte
	Content interface{}
	From    string
}

func CreateId(uid, toUid string) string {
	return uid + "->" + toUid
}
