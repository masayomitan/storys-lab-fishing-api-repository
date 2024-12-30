package validator

import (
	"fmt"

	// "github.com/go-playground/validator/v10"
)

func GenerateCustomMessage(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%sは必須です。", field)
	case "max":
		return fmt.Sprintf("%sは%s文字以内で入力してください。", field, param)
	case "min":
		return fmt.Sprintf("%sは%s文字以上で入力してください。", field, param)
	default:
		return fmt.Sprintf("%sの入力が無効です。", field)
	}
}
