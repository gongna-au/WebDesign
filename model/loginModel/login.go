package loginModel

import (
	"github.com/WebDesign/pkg/validaors"
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
		"phone": []string{"required", "digits:11"},
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

	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validaors.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
