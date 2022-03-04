package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Validators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
	}
}

func nameNotNullAndAdmin(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value != "" && "admin" != value
}
