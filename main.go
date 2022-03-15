/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/15/2022
    @Note:       main函数
**/

package main

import (
	"chat_demo/core"
	"chat_demo/dao/mongodb"
	"chat_demo/dao/mysql"
	"chat_demo/dao/redis"
	"chat_demo/pkg/initialize"

	"chat_demo/global"
)

func main() {
	// 写入配置
	// 先写入配置文件viper
	global.Chat_VP = core.Viper("./config.yaml")
	global.Chat_LOG = core.Zap()
	// 初始化redis
	redis.Redis()
	// 初始化mysql
	mysql.Mysql()
	// 初始化mongodb
	mongodb.MongoDb()

	if global.Chat_REDIS != nil {
		defer redis.Close()
	}
	if global.Chat_MYSQL != nil {
		initialize.RegisterTables(global.Chat_MYSQL)
		defer mysql.Close()
	}
	if global.Chat_MONGODB != nil {
		defer mongodb.Close()
	}
	// 启动管道监听
	go core.RunMonitor()
	// 启动服务
	core.RunServer()
}
