/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       Redis连接/关闭
**/

package redis

import (
	"chat_demo/global"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.Chat_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisCfg.Host,
			redisCfg.Port,
		),
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
		PoolSize: redisCfg.PoolSize,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Chat_LOG.Error("[Redis] connect ping failed, err:", zap.Error(err))
	} else {
		global.Chat_LOG.Info("[Redis] connect ping success")
		global.Chat_REDIS = client
	}
}

func Close() {
	_ = global.Chat_REDIS.Close()
}
