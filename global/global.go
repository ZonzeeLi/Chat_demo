/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       全局变量
**/

package global

import (
	"chat_demo/conf"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Chat_VP      *viper.Viper
	Chat_REDIS   *redis.Client
	Chat_LOG     *zap.Logger
	Chat_CONFIG  conf.Server
	Chat_MONGODB *mongo.Client
	Chat_MYSQL   *gorm.DB
)
