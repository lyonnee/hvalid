package complex

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// ConvertValidator 转换验证器结构体
type ConvertValidator[T, U any] struct {
	FieldName string // 字段名称
}

// NewConvertValidator 创建转换验证器
func NewConvertValidator[T, U any](fieldName string) *ConvertValidator[T, U] {
	return &ConvertValidator[T, U]{
		FieldName: fieldName,
	}
}

// Convert 转换值后验证
func (v *ConvertValidator[T, U]) Convert(convert func(T) U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		converted := convert(value)
		return validator(converted)
	})
}

// ConvertWithError 转换值后验证（支持错误处理）
func (v *ConvertValidator[T, U]) ConvertWithError(convert func(T) (U, error), validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		converted, err := convert(value)
		if err != nil {
			return fmt.Errorf("conversion failed: %w", err)
		}
		return validator(converted)
	})
}

// ConvertSlice 转换切片中的每个元素后验证
func (v *ConvertValidator[T, U]) ConvertSlice(convert func(T) U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(values []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for i, value := range values {
			converted := convert(value)
			if err := validator(converted); err != nil {
				validationErr.AddError(fmt.Sprintf("element[%d]: %v", i, err))
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// ConvertMap 转换映射中的值后验证
func (v *ConvertValidator[T, U]) ConvertMap(convert func(T) U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[map[string]T] {
	return hvalid.ValidatorFunc[map[string]T](func(values map[string]T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for key, value := range values {
			converted := convert(value)
			if err := validator(converted); err != nil {
				validationErr.AddError(fmt.Sprintf("key[%s]: %v", key, err))
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// ConvertWithDefault 使用默认值转换后验证
func (v *ConvertValidator[T, U]) ConvertWithDefault(convert func(T) (U, error), defaultValue U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		converted, err := convert(value)
		if err != nil {
			converted = defaultValue
		}
		return validator(converted)
	})
}
