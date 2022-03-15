/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/11/2022
    @UpdateDate: 3/11/2022
    @Note:       用户注册数据库绑定结构
**/

package register

import "gorm.io/gorm"

type UserRegisterInfo struct {
	gorm.Model
	UserName       string `json:"user_name" gorm:"comment:用户名"`
	PasswordDigest string `json:"password_digest" gorm:"comment:密码"`
	Email          string `json:"email" gorm:"comment:邮箱"`
	Avatar         string `json:"avatar" gorm:"comment:头像"`
	Phone          string `json:"phone" gorm:"comment:电话"`
	Status         string `json:"status" gorm:"comment:状态"`
}

const Active string = "active"
