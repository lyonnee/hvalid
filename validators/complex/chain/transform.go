package complex

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// TransformValidator 转换验证器结构体
type TransformValidator[T, U any] struct {
	FieldName string // 字段名称
}

// NewTransformValidator 创建转换验证器
func NewTransformValidator[T, U any](fieldName string) *TransformValidator[T, U] {
	return &TransformValidator[T, U]{
		FieldName: fieldName,
	}
}

// Transform 转换值后验证
func (v *TransformValidator[T, U]) Transform(transform func(T) U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		transformed := transform(value)
		return validator(transformed)
	})
}

// TransformWithError 转换值后验证（支持错误处理）
func (v *TransformValidator[T, U]) TransformWithError(transform func(T) (U, error), validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		transformed, err := transform(value)
		if err != nil {
			return fmt.Errorf("transform failed: %w", err)
		}
		return validator(transformed)
	})
}

// Map 对切片中的每个元素进行转换和验证
func (v *TransformValidator[T, U]) Map(transform func(T) U, validator hvalid.ValidatorFunc[U]) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(values []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for i, value := range values {
			transformed := transform(value)
			if err := validator(transformed); err != nil {
				validationErr.AddError(fmt.Sprintf("element[%d]: %v", i, err))
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// Filter 过滤并验证元素
func (v *TransformValidator[T, U]) Filter(filter func(T) bool, validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[[]T] {
	return hvalid.ValidatorFunc[[]T](func(values []T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for i, value := range values {
			if filter(value) {
				if err := validator(value); err != nil {
					validationErr.AddError(fmt.Sprintf("element[%d]: %v", i, err))
				}
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}
