package primitive

import (
	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNotTrue  = "must be true"
	ErrNotFalse = "must be false"
)

// BooleanValidator 布尔值验证器结构体
type BooleanValidator struct {
	FieldName string // 字段名称
}

// NewBooleanValidator 创建布尔值验证器
func NewBooleanValidator(fieldName string) *BooleanValidator {
	return &BooleanValidator{
		FieldName: fieldName,
	}
}

// IsTrue 验证值必须为 true
func (v *BooleanValidator) IsTrue() hvalid.ValidatorFunc[bool] {
	return hvalid.ValidatorFunc[bool](func(value bool) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !value {
			validationErr.AddError(ErrNotTrue)
			return validationErr
		}
		return nil
	})
}

// IsFalse 验证值必须为 false
func (v *BooleanValidator) IsFalse() hvalid.ValidatorFunc[bool] {
	return hvalid.ValidatorFunc[bool](func(value bool) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if value {
			validationErr.AddError(ErrNotFalse)
			return validationErr
		}
		return nil
	})
}
