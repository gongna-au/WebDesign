package login

import (
	"github.com/WebDesign/handler"
	"github.com/WebDesign/model/loginModel"
	"github.com/WebDesign/model/requests"
	"github.com/WebDesign/model/response"
	"github.com/WebDesign/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	handler.BaseAPIController
}

// LoginByPhone 手机登录
func LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := loginModel.LoginByPhoneRequest{}

	if ok := handler.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := AttemptLoginByPhone(request.Phone, request.Password)

	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}

}
