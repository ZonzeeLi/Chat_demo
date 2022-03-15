/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       Mysql配置结构
**/

package conf

import "fmt"

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"` // 高级配置
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"` // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`    // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
		m.Config,
	)
	//return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
