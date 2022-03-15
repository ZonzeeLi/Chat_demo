/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/13/2022
    @Note:       服务监听
**/

package core

import (
	"chat_demo/global"
	routers "chat_demo/routers"
	"chat_demo/service/chat"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunServer() {
	router := routers.Routers()
	address := fmt.Sprintf(":%d", global.Chat_CONFIG.System.Port)
	s := initServer(address, router)
	global.Chat_LOG.Info("[Server] run success on ", zap.String("address", address))
	global.Chat_LOG.Error(s.ListenAndServe().Error())
}

func RunMonitor() {
	global.Chat_LOG.Info("[Monitor] run start...")
	chat.Monitor()
}
