package complex

import (
	"github.com/lyonnee/hvalid"
)

// ChainValidator 链式验证器结构体
type ChainValidator[T any] struct {
	FieldName  string // 字段名称
	validators []hvalid.ValidatorFunc[T]
}

// NewChainValidator 创建链式验证器
func NewChainValidator[T any](fieldName string) *ChainValidator[T] {
	return &ChainValidator[T]{
		FieldName:  fieldName,
		validators: make([]hvalid.ValidatorFunc[T], 0),
	}
}

// Add 添加验证器到链中
func (v *ChainValidator[T]) Add(validator hvalid.ValidatorFunc[T]) *ChainValidator[T] {
	v.validators = append(v.validators, validator)
	return v
}

// Validate 执行链式验证
func (v *ChainValidator[T]) Validate(value T) error {
	validationErr := hvalid.NewValidationError(v.FieldName)

	for _, validator := range v.validators {
		if err := validator(value); err != nil {
			validationErr.AddError(err.Error())
		}
	}

	if validationErr.HasError() {
		return validationErr
	}
	return nil
}

// ValidateFirstError 执行链式验证，返回第一个错误
func (v *ChainValidator[T]) ValidateFirstError(value T) error {
	for _, validator := range v.validators {
		if err := validator(value); err != nil {
			return err
		}
	}
	return nil
}

// ValidateAllErrors 执行链式验证，返回所有错误
func (v *ChainValidator[T]) ValidateAllErrors(value T) error {
	validationErr := hvalid.NewValidationError(v.FieldName)

	for _, validator := range v.validators {
		if err := validator(value); err != nil {
			validationErr.AddError(err.Error())
		}
	}

	if validationErr.HasError() {
		return validationErr
	}
	return nil
}

// Clear 清除所有验证器
func (v *ChainValidator[T]) Clear() *ChainValidator[T] {
	v.validators = make([]hvalid.ValidatorFunc[T], 0)
	return v
}

// Length 获取验证器链的长度
func (v *ChainValidator[T]) Length() int {
	return len(v.validators)
}
