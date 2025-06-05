package complex

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrAllValidatorsFailed = "all validators failed"
	ErrAnyValidatorFailed  = "any validator failed"
)

// LogicValidator 逻辑组合验证器结构体
type LogicValidator[T any] struct {
	FieldName string // 字段名称
}

// NewLogicValidator 创建逻辑组合验证器
func NewLogicValidator[T any](fieldName string) *LogicValidator[T] {
	return &LogicValidator[T]{
		FieldName: fieldName,
	}
}

// All 所有验证器都必须通过
func (v *LogicValidator[T]) All(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, validator := range validators {
			if err := validator(value); err != nil {
				validationErr.AddError(err.Error())
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// Any 任意一个验证器通过即可
func (v *LogicValidator[T]) Any(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, validator := range validators {
			validatorErr := validator(value)
			if validatorErr == nil {
				return nil
			}
			validationErr.AddError(validatorErr.Error())
		}

		return validationErr
	})
}

// None 所有验证器都必须失败
func (v *LogicValidator[T]) None(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, validator := range validators {
			if err := validator(value); err == nil {
				validationErr.AddError(fmt.Sprintf("validator should fail: %v", validator))
			}
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// Not 验证器必须失败
func (v *LogicValidator[T]) Not(validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if err := validator(value); err == nil {
			return fmt.Errorf("validator should fail: %v", validator)
		}
		return nil
	})
}
