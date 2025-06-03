package primitive

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrSliceTooShort = "length must be at least %d"
	ErrSliceTooLong  = "length must be at most %d"
	ErrSliceEmpty    = "must not be empty"
	ErrSliceNotEmpty = "must be empty"
	ErrSliceContains = "must contain the element"
)

// SliceValidator 切片验证器结构体
type SliceValidator[T any] struct {
	FieldName string // 字段名称
}

// NewSliceValidator 创建切片验证器
func NewSliceValidator[T any](fieldName string) *SliceValidator[T] {
	return &SliceValidator[T]{
		FieldName: fieldName,
	}
}

// MinLen 验证最小长度
func (v *SliceValidator[T]) MinLen(minLen int) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(slice []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(slice) < minLen {
			validationErr.AddError(fmt.Sprintf(ErrSliceTooShort, minLen))
			return validationErr
		}
		return nil
	})
}

// MaxLen 验证最大长度
func (v *SliceValidator[T]) MaxLen(maxLen int) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(slice []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(slice) > maxLen {
			validationErr.AddError(fmt.Sprintf(ErrSliceTooLong, maxLen))
			return validationErr
		}
		return nil
	})
}

// NotEmpty 验证不能为空
func (v *SliceValidator[T]) NotEmpty() hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(slice []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(slice) == 0 {
			validationErr.AddError(ErrSliceEmpty)
			return validationErr
		}
		return nil
	})
}

// Empty 验证必须为空
func (v *SliceValidator[T]) Empty() hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(slice []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(slice) > 0 {
			validationErr.AddError(ErrSliceNotEmpty)
			return validationErr
		}
		return nil
	})
}

// Contains 验证是否包含指定元素
func (v *SliceValidator[T]) Contains(element T) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(slice []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, item := range slice {
			if fmt.Sprintf("%v", item) == fmt.Sprintf("%v", element) {
				return nil
			}
		}
		validationErr.AddError(ErrSliceContains)
		return validationErr
	})
}
