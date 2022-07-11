package login

import (
	"errors"

	//"github.com/WebDesign/log"
	"github.com/WebDesign/model/user"
	//"github.com/gin-gonic/gin"
)

// LoginByPhone 登录指定用户
func AttemptLoginByPhone(phone string, password string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	if userModel.Password != password {
		return user.User{}, errors.New("账号和密码不正确")
	}

	return userModel, nil
}
