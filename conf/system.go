/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/11/2022
    @Note:       系统服务配置结构
**/

package conf

type System struct {
	Name   string `mapstructrue:"name" json:"name" yaml:"name"`
	Port   int    `mapstructrue:"port" json:"port" yaml:"port"`
	Author string `mapstructrue:"author" json:"author" yaml:"author"`
}
