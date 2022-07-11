package requests

import (
	//db "github.com/WebDesign/database"
	//"github.com/WebDesign/model/user"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`
	Password string `json:"password,omitempty" valid:"password"`
}

// LoginByPhone 验证表单，返回长度等于零即通过

func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"password": []string{"required"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码为必填项，参数名称 password",
		},
	}

	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
