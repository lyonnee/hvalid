package primitive

import (
	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrTextTooShort = "value length too short"
	ErrTextTooLong  = "value length too long"
)

// TextValidator 文本验证器结构体
type TextValidator[T string | []byte] struct {
	FieldName string // 字段名称
}

// NewTextValidator 创建文本验证器
func NewTextValidator[T string | []byte](fieldName string) *TextValidator[T] {
	return &TextValidator[T]{
		FieldName: fieldName,
	}
}

// MinLen 验证最小长度
func (v *TextValidator[T]) MinLen(minLen int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(field T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		l := len(field)
		if l < minLen {
			validationErr.AddError(ErrTextTooShort)
			return validationErr
		}
		return nil
	})
}

// MaxLen 验证最大长度
func (v *TextValidator[T]) MaxLen(maxLen int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(field T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		l := len(field)
		if l > maxLen {
			validationErr.AddError(ErrTextTooLong)
			return validationErr
		}
		return nil
	})
}
