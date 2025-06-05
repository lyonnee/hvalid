package complex

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// DependencyValidator 依赖验证器结构体
type DependencyValidator[T any] struct {
	FieldName string // 字段名称
}

// NewDependencyValidator 创建依赖验证器
func NewDependencyValidator[T any](fieldName string) *DependencyValidator[T] {
	return &DependencyValidator[T]{
		FieldName: fieldName,
	}
}

// DependsOn 验证器依赖于另一个验证器的结果
func (v *DependencyValidator[T]) DependsOn(dependency hvalid.ValidatorFunc[T], validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		// 首先验证依赖
		if err := dependency(value); err != nil {
			return fmt.Errorf("dependency validation failed: %v", err)
		}

		// 依赖验证通过后，执行主验证
		return validator(value)
	})
}

// DependsOnAll 验证器依赖于多个验证器的结果
func (v *DependencyValidator[T]) DependsOnAll(dependencies []hvalid.ValidatorFunc[T], validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		// 验证所有依赖
		for i, dependency := range dependencies {
			if err := dependency(value); err != nil {
				validationErr.AddError(fmt.Sprintf("dependency[%d] validation failed: %v", i, err))
			}
		}

		// 如果有依赖验证失败，直接返回
		if validationErr.HasError() {
			return validationErr
		}

		// 所有依赖验证通过后，执行主验证
		return validator(value)
	})
}

// DependsOnAny 验证器依赖于任意一个验证器的结果
func (v *DependencyValidator[T]) DependsOnAny(dependencies []hvalid.ValidatorFunc[T], validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)
		anySuccess := false

		// 验证所有依赖
		for i, dependency := range dependencies {
			if err := dependency(value); err == nil {
				anySuccess = true
				break
			} else {
				validationErr.AddError(fmt.Sprintf("dependency[%d] validation failed: %v", i, err))
			}
		}

		// 如果没有任何依赖验证通过，返回错误
		if !anySuccess {
			return validationErr
		}

		// 至少有一个依赖验证通过后，执行主验证
		return validator(value)
	})
}

// DependsOnCondition 验证器依赖于条件的结果
func (v *DependencyValidator[T]) DependsOnCondition(condition func(T) bool, validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		// 检查条件
		if !condition(value) {
			return fmt.Errorf("dependency condition not met")
		}

		// 条件满足后，执行主验证
		return validator(value)
	})
}
