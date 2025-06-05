package complex

import (
	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrConditionNotMet = "condition not met: %s"
)

// ConditionValidator 条件验证器结构体
type ConditionValidator[T any] struct {
	FieldName string // 字段名称
}

// NewConditionValidator 创建条件验证器
func NewConditionValidator[T any](fieldName string) *ConditionValidator[T] {
	return &ConditionValidator[T]{
		FieldName: fieldName,
	}
}

// When 当条件满足时执行验证
func (v *ConditionValidator[T]) When(condition bool, validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if condition {
			return validator(value)
		}
		return nil
	})
}

// Unless 当条件不满足时执行验证
func (v *ConditionValidator[T]) Unless(condition bool, validator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return v.When(!condition, validator)
}

// If 根据条件选择不同的验证器
func (v *ConditionValidator[T]) If(condition bool, ifValidator, elseValidator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		if condition {
			return ifValidator(value)
		}
		return elseValidator(value)
	})
}

// Switch 根据条件选择不同的验证器
func (v *ConditionValidator[T]) Switch(cases map[bool]hvalid.ValidatorFunc[T], defaultValidator hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		for condition, validator := range cases {
			if condition {
				return validator(value)
			}
		}
		return defaultValidator(value)
	})
}
