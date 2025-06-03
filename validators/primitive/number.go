package primitive

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNumberTooSmall = "must be greater than or equal to %v"
	ErrNumberTooBig   = "must be less than or equal to %v"
)

// NumberValidator 数字验证器结构体
type NumberValidator[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64] struct {
	FieldName string // 字段名称
}

// NewNumberValidator 创建数字验证器
func NewNumberValidator[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](fieldName string) *NumberValidator[T] {
	return &NumberValidator[T]{
		FieldName: fieldName,
	}
}

// Min 验证最小值
func (v *NumberValidator[T]) Min(min T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(num T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if num < min {
			validationErr.AddError(fmt.Sprintf(ErrNumberTooSmall, min))
			return validationErr
		}
		return nil
	})
}

// Max 验证最大值
func (v *NumberValidator[T]) Max(max T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(num T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if num > max {
			validationErr.AddError(fmt.Sprintf(ErrNumberTooBig, max))
			return validationErr
		}
		return nil
	})
}
