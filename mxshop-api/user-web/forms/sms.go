package forms

type SendSmsForm struct {
	Mobile string `from:"mobile" json:"mobile" binding:"required,mobile"` //电话号码有什么规律可寻，需要自定义validator
	Type   uint   `from:"type" json:"type" binding:"required,oneof=1 2"`  //1表示注册，2表示登录 需要使用type区别，验证码的业务类别
}
