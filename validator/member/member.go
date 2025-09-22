package member

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(fl validator.FieldLevel) bool {
	// 获取字段值
	name := fl.Field().String()

	// 判断是否为 "admin"
	if name == "admin" {
		return false
	}

	return true
}
