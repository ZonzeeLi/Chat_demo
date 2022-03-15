## IM 即时聊天Demo

&emsp;&emsp;在[该作者项目](https://github.com/CocaineCong/gin-chat-demo)基础上，将一些功能进行完善，还有一些错误的地方以及未开发的部分补充完整，修改了整体项目结构和一些工具包。

## 项目介绍

&emsp;&emsp;该Demo是基于Go+Gin+WebSocket+Mysql+Redis+MongoDb完成的一个功能简单，二次开发、结构性非常强的一个双人聊天demo(可扩展成群聊)。

### 项目功能

1. 使用Mysql完成用户注册验证。
2. Gin框架升级WebSocket协议，用Redis做缓存实现双人在线聊天，可流传输文本等。
3. 用MongoDb存储聊天记录信息，并对双人聊天的双方进行发送和接收的消息结构储存。
4. 消息存储包含已读/未读状态，在接收消息可返回历史聊天记录和未读信息，并且在MongoDb中更新为已读状态，刷新过期时间。

## 项目结构

```项目目录结构
├─api					(api层)
│  ├─chat				(ChatApi管理)
│  └─register			(RegisterApi管理)
├─conf					(全局配置层)
├─core					(核心文件)
├─dao					(数据库层)
│  ├─mongodb			(MongoDb连接/关闭)
│  ├─mysql				(MySql连接/关闭)
│  └─redis				(Redis连接/关闭)
├─global				(全局对象)
├─log					(日志，日志文件自动生成)
├─model					(模型层)
│  ├─chat				(聊天结构)
│  ├─register			(用户注册结构)
│  └─response			(响应结构)
├─pkg					(工具包)
│  ├─initialize			(初始化)
│  └─utils				(工具)
├─routers				(路由层，可继续扩展成路由组管理，类似Api和Service)
└─service				(服务层)
	├─chat				(ChatService管理)
    └─register			(RegisterService管理)

```

## 项目使用

### 配置

&emsp;&emsp;只要修改config.yaml的数据库配置和服务配置即可，日志使用的是zap日志，也可以自定义zap日志核心文件或配置zap日志参数，数据库使用的是gorm，gorm的日志追踪也是自定义，可以自行修改。

### 启动服务

```go
go run main.go
```

### 聊天Api使用

#### 建立连接

&emsp;&emsp;使用的是postman进行websocket请求

[连接](picture/连接.png)

#### 离线发送消息

&emsp;&emsp;发送方连接到服务器并成功注册到用户管理，而接收方没有，返回提示"对方不在线"。

[离线发送](./picture/离线发送.png)

#### 在线发送消息

&emsp;&emsp;双方都连接到服务器并成功注册到用户管理，返回提示"对方在线"。

[在线发送](./picture/在线发送.png)

#### 查询历史记录

&emsp;&emsp;按照时间排序，查询双方历史记录(代码中写入为10条历史消息)。

[查询历史记录](./picture/查询历史记录.png)

#### 查询未读消息

&emsp;&emsp;如果在线则默认为已读(逻辑可以根据实际进行修改)，离线默认为未读，如果一直在线则没有未读消息，返回提示"没有未读消息"，如果有离线情况(即未读情况)，返回双方所有未读消息，并更新己方未读消息在数据库中的状态为已读，且更新过期时间。

[未读消息](./picture/未读消息.png)

[查询未读消息](./picture/查询未读消息.png)

## 总结

&emsp;&emsp;该项目整体结构是参考了[GVA](https://github.com/flipped-aurora/gin-vue-admin) 的结构来做的，由于其优秀的封装思路和设计，将不同层内的不同api或对象等封装为group使用，所以该项目同样也具有很好的结构性，二次开发比较容易。同时很多地方做的都比较简单，比如登陆和注册验证并没有加密、用户信息等内容并不完善、数据传输单一，这些地方都可以根据实际场景修改内部逻辑，另外Websocket传输最好直接用[]byte来做，该项目为了调试观察方便，做了转换，和前端对接应该直接传输[]byte，群聊的设计是再开一个用户管理，用[]Client来设计。