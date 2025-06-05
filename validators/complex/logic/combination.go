package complex

import (
	"reflect"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNotEqual = "the two values are not equal"
	ErrEmpty    = "the value is empty"
)

// CombinationValidator 组合验证器结构体
type CombinationValidator[T any] struct {
	FieldName string // 字段名称
}

// NewCombinationValidator 创建组合验证器
func NewCombinationValidator[T any](fieldName string) *CombinationValidator[T] {
	return &CombinationValidator[T]{
		FieldName: fieldName,
	}
}

// Eq 验证值是否相等
func (v *CombinationValidator[T]) Eq(comparData T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !reflect.DeepEqual(data, comparData) {
			validationErr.AddError(ErrNotEqual)
			return validationErr
		}
		return nil
	})
}

// Required 验证值是否非空
func (v *CombinationValidator[T]) Required() hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		var t interface{} = data
		if t == nil {
			validationErr.AddError(ErrEmpty)
			return validationErr
		}

		rv := reflect.ValueOf(t)
		if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface || rv.Kind() == reflect.Func) && rv.IsNil() {
			validationErr.AddError(ErrEmpty)
			return validationErr
		}
		return nil
	})
}
