package primitive

import (
	"bytes"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrBytesNotContains = "must contain the sub byte slice"
)

// BytesValidator 字节切片验证器结构体
type BytesValidator struct {
	FieldName string // 字段名称
}

// NewBytesValidator 创建字节切片验证器
func NewBytesValidator(fieldName string) *BytesValidator {
	return &BytesValidator{
		FieldName: fieldName,
	}
}

// ContainsBytes 验证是否包含子字节切片
func (v *BytesValidator) ContainsBytes(subslice []byte) hvalid.ValidatorFunc[[]byte] {
	return hvalid.ValidatorFunc[[]byte](func(field []byte) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		ok := bytes.Contains(field, subslice)
		if !ok {
			validationErr.AddError(ErrBytesNotContains)
			return validationErr
		}

		return nil
	})
}
