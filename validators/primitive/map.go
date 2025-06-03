package primitive

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrMapTooShort = "must have at least %d entries"
	ErrMapTooLong  = "must have at most %d entries"
	ErrMapEmpty    = "must not be empty"
	ErrMapNotEmpty = "must be empty"
	ErrMapHasKey   = "must not contain key: %v"
	ErrMapNoKey    = "must contain key: %v"
)

// MapValidator Map验证器结构体
type MapValidator[K comparable, V any] struct {
	FieldName string // 字段名称
}

// NewMapValidator 创建Map验证器
func NewMapValidator[K comparable, V any](fieldName string) *MapValidator[K, V] {
	return &MapValidator[K, V]{
		FieldName: fieldName,
	}
}

// MinSize 验证最小大小
func (v *MapValidator[K, V]) MinSize(minSize int) hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(m) < minSize {
			validationErr.AddError(fmt.Sprintf(ErrMapTooShort, minSize))
			return validationErr
		}
		return nil
	})
}

// MaxSize 验证最大大小
func (v *MapValidator[K, V]) MaxSize(maxSize int) hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(m) > maxSize {
			validationErr.AddError(fmt.Sprintf(ErrMapTooLong, maxSize))
			return validationErr
		}
		return nil
	})
}

// NotEmpty 验证不能为空
func (v *MapValidator[K, V]) NotEmpty() hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(m) == 0 {
			validationErr.AddError(ErrMapEmpty)
			return validationErr
		}
		return nil
	})
}

// Empty 验证必须为空
func (v *MapValidator[K, V]) Empty() hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if len(m) > 0 {
			validationErr.AddError(ErrMapNotEmpty)
			return validationErr
		}
		return nil
	})
}

// HasKey 验证是否包含指定键
func (v *MapValidator[K, V]) HasKey(key K) hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if _, exists := m[key]; !exists {
			validationErr.AddError(fmt.Sprintf(ErrMapNoKey, key))
			return validationErr
		}
		return nil
	})
}

// NoKey 验证不能包含指定键
func (v *MapValidator[K, V]) NoKey(key K) hvalid.ValidatorFunc[map[K]V] {
	return hvalid.ValidatorFunc[map[K]V](func(m map[K]V) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if _, exists := m[key]; exists {
			validationErr.AddError(fmt.Sprintf(ErrMapHasKey, key))
			return validationErr
		}
		return nil
	})
}
