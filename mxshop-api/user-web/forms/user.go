package forms

type PassWordLoginForm struct {
	Mobile    string `from:"mobile" json:"mobile" binding:"required,mobile"` //电话号码有什么规律可寻，需要自定义validator
	Password  string `from:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `from:"captcha" json:"captcha" binding:"required,min=4,max=4"` //验证码
	CaptchaId string `from:"captcha_id" json:"captcha_id" binding:"required"`       //验证码id
}

type RegisterForm struct {
	Mobile   string `from:"mobile" json:"mobile" binding:"required,mobile"` //电话号码有什么规律可寻，需要自定义validator
	Password string `from:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `from:"code" json:"code" binding:"required,min=5,max=5"`
}
