/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       Redis配置结构
**/

package conf

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                      // redis的哪个数据库
	Host     string `mapstructure:"host" json:"host" yaml:"host"`                // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`                // 端口
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // 密码
	PoolSize int    `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"` // 连接池
}
