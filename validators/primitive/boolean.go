package primitive

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNotTrue  = "must be true"
	ErrNotFalse = "must be false"
)

// IsTrue 验证值必须为 true
func IsTrue() hvalid.ValidatorFunc[bool] {
	return hvalid.ValidatorFunc[bool](func(value bool) error {
		if !value {
			return errors.New(ErrNotTrue)
		}
		return nil
	})
}

// IsFalse 验证值必须为 false
func IsFalse() hvalid.ValidatorFunc[bool] {
	return hvalid.ValidatorFunc[bool](func(value bool) error {
		if value {
			return errors.New(ErrNotFalse)
		}
		return nil
	})
}
