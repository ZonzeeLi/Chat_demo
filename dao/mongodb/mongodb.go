/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       MongoDb连接/关闭
**/

package mongodb

import (
	"chat_demo/global"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func MongoDb() {
	uri := fmt.Sprintf("mongodb://%s:%d",
		global.Chat_CONFIG.MongoDb.Host,
		global.Chat_CONFIG.MongoDb.Port)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(uint64(global.Chat_CONFIG.MongoDb.Poolsize))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.Chat_LOG.Error("[MongoDB] connect failed, err:", zap.Error(err))
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		global.Chat_LOG.Error("[MongoDB] ping failed, err:", zap.Error(err))
	} else {
		global.Chat_LOG.Info("[MongoDB] connect ping success")
		global.Chat_MONGODB = client
	}
}

func Close() {
	_ = global.Chat_MONGODB.Disconnect(context.TODO())
}
