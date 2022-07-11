package validaors

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	return errs
}
