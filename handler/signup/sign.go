package signup

import (
	//"fmt"

	"github.com/WebDesign/handler"
	"github.com/WebDesign/model/requests"
	"github.com/WebDesign/model/response"
	"github.com/WebDesign/model/user"
	"github.com/WebDesign/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	handler.BaseAPIController
}

// SignupUsingPhone 使用手机和密码进行注册
func SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	//requests.SignupUsingPhone 验证函数
	if ok := handler.Validate(c, &request, requests.SignupUsingPhone); !ok {

		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

// IsPhoneExist 检测手机号是否被注册
func IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}

	if ok := handler.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": requests.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := handler.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": requests.IsEmailExist(request.Email),
	})
}
