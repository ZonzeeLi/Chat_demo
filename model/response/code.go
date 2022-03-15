/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/15/2022
    @Note:       状态码信息
**/

package response

type ResCode int64

const (
	CodeSuccess               ResCode = 200
	CodeUpdatePasswordSuccess ResCode = 201
	CodeUserNotExit           ResCode = 202
	CodeUserExit              ResCode = 203
	CodeInvalidParams         ResCode = 400

	CodeServerError ResCode = 500

	CodeDatabaseError  ResCode = 1001
	CodeWSSuccess      ResCode = 1200
	CodeWSSuccessMsg   ResCode = 1201
	CodeWSEnd          ResCode = 1203
	CodeWSOnlineReply  ResCode = 1204
	CodeWSOfflineReply ResCode = 1205
	CodeWSLimit        ResCode = 1206
	CodeWSNoHistory    ResCode = 1207
	CodeWsNoUnread     ResCode = 1208
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:               "成功",
	CodeUpdatePasswordSuccess: "修改成功",
	CodeUserNotExit:           "用户不存在",
	CodeUserExit:              "用户已存在",
	CodeInvalidParams:         "请求参数错误",

	CodeServerError:    "服务错误",
	CodeDatabaseError:  "数据库操作失败",
	CodeWSSuccess:      "连接成功",
	CodeWSSuccessMsg:   "解析content内容信息",
	CodeWSEnd:          "断开连接",
	CodeWSOnlineReply:  "对方在线",
	CodeWSOfflineReply: "对方离线",
	CodeWSLimit:        "请求收到限制",
	CodeWSNoHistory:    "没有历史记录",
	CodeWsNoUnread:     "没有未读消息",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerError]
	}
	return msg
}
