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

// Min 验证数值是否大于等于最小值
func Min[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](min T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if value < min {
			return fmt.Errorf("must be greater than or equal to %v", min)
		}
		return nil
	})
}

// Max 验证数值是否小于等于最大值
func Max[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](max T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if value > max {
			return fmt.Errorf("must be less than or equal to %v", max)
		}
		return nil
	})
}

// Range 验证数值是否在指定范围内
func Range[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](min, max T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if value < min || value > max {
			return fmt.Errorf("must be between %v and %v", min, max)
		}
		return nil
	})
}

// Positive 验证数值是否为正数
func Positive[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64]() hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if value <= 0 {
			return fmt.Errorf("must be positive")
		}
		return nil
	})
}

// Negative 验证数值是否为负数
func Negative[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64]() hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if value >= 0 {
			return fmt.Errorf("must be negative")
		}
		return nil
	})
}
