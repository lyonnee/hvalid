package complex

import (
	"errors"
	"reflect"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNotEqual = "the two values are not equal"
	ErrEmpty    = "the value is empty"
)

func Eq[T comparable](comparData T) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		if data != comparData {
			return errors.New(ErrNotEqual)
		}
		return nil
	})
}

func Required[T any]() hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		var t interface{} = data
		if t == nil {
			return errors.New(ErrEmpty)
		}

		rv := reflect.ValueOf(t)
		if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface || rv.Kind() == reflect.Func) && rv.IsNil() {
			return errors.New(ErrEmpty)
		}
		return nil
	})
}
