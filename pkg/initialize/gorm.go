/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/11/2022
    @UpdateDate: 3/11/2022
    @Note:       Gorm数据库表绑定
**/

package initialize

import (
	"chat_demo/global"
	"chat_demo/model/register"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterTables(db *gorm.DB) {

	err := db.AutoMigrate(
		register.UserRegisterInfo{},
	)

	if err != nil {
		global.Chat_LOG.Error("[Init] register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Chat_LOG.Info("[Init] register table success")
}
