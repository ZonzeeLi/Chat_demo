/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       MongoDb配置结构
**/

package conf

type MongoDb struct {
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Poolsize int    `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"`
}
