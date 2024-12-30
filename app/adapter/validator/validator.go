package validator

import (
	"fmt"
	// "strconv"
	"strings"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

// NewValidator は新しいバリデーションインスタンスを作成します。
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct は構造体に対するバリデーションを行います。
func (v *Validator) ValidateStruct(s interface{}) error {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	// バリデーションエラーをカスタムメッセージに変換
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			// タグ（例: max=20）からパラメータを分解
			param := e.Param()

			// フィールド名、タグ名、パラメータを基にメッセージ生成
			customMessage := GenerateCustomMessage(e.Field(), e.Tag(), param)
			errorMessages = append(errorMessages, customMessage)
		}
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}

	return err
}

// // カスタムバリデーションの例
// func (v *Validator) RegisterCustomValidations() {
// 	// カスタムバリデーション "is-even" を登録
// 	v.validate.RegisterValidation("is-even", func(fl validator.FieldLevel) bool {
// 		num, ok := fl.Field().Interface().(int)
// 		if !ok {
// 			return false
// 		}
// 		return num%2 == 0
// 	})
// }
