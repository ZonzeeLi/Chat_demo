/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/13/2022
    @Note:       Mysql连接/关闭
**/

package mysql

import (
	"chat_demo/global"
	"chat_demo/pkg/initialize"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	m := global.Chat_CONFIG.Mysql
	if m.Dbname == "" {
		global.Chat_LOG.Error("[Mysql] DBname is empty, please enter. ")
		return
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), initialize.Gorm.Config())
	if err != nil {
		global.Chat_LOG.Error("[Mysql] connect failed, err:", zap.Error(err))
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		global.Chat_LOG.Error("[Mysql] ping failed, err:", zap.Error(err))
		return
	}
	global.Chat_LOG.Info("[Mysql] connect ping success")
	global.Chat_MYSQL = db

}

func Close() {
	db, _ := global.Chat_MYSQL.DB()
	db.Close()
}
