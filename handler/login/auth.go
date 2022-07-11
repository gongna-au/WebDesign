package login

import (
	"errors"

	"github.com/WebDesign/log"
	"github.com/WebDesign/model/user"
	"github.com/gin-gonic/gin"
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

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		log.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
