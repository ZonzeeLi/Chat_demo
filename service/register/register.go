/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/11/2022
    @UpdateDate: 3/14/2022
    @Note:       Service层 RegisterService服务管理
**/

package register

import (
	"chat_demo/global"
	"chat_demo/model/register"
	"chat_demo/model/response"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

type RegisterService struct{}

func (registerService *RegisterService) Register(user register.UserRegisterRequest) (rescode response.ResCode, resmsg string) {
	var userinfo register.UserRegisterInfo
	if err := global.Chat_MYSQL.Where("user_name = ?", user.UserName).First(&userinfo).Error; err == gorm.ErrRecordNotFound {
		// 可自行对用户密码进行加密
		userinfo = register.UserRegisterInfo{
			UserName:       user.UserName,
			PasswordDigest: user.Password,
			Status:         register.Active, // 默认初始状态
		}
		userinfo.Avatar = "https://lmg.jj20.com/up/allimg/tx27/59111103098507.jpg" // 默认初始头像
		err = global.Chat_MYSQL.Create(&userinfo).Error
		if err != nil {
			global.Chat_LOG.Error("[Register] User create failed, err:", zap.Error(err))
			rescode = response.CodeDatabaseError
		} else {
			rescode = response.CodeSuccess
		}
	} else {
		rescode = response.CodeUserExit
		resmsg = response.CodeUserExit.Msg() // 可以修改为error.New
	}
	return
}
