package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

//ValidatorMobile 自定义validator
func ValidatorMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式准标准库：regexp
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
