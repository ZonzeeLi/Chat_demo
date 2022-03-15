/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       全局配置结构
**/

package conf

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	MongoDb MongoDb `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
}
