package composite

import (
	"errors"
	"reflect"

	"github.com/lyonnee/hvalid"
)

func Eq[T comparable](comparData T, errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		if data != comparData {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("the two values are not equal")
		}

		return nil
	})
}

func Required[T any](errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(data T) error {
		err := errors.New("the value is empty")
		if len(errMsg) > 0 {
			err = errors.New(errMsg[0])
		}

		var t interface{} = data
		if t == nil {
			return err
		}

		rv := reflect.ValueOf(t)
		if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface || rv.Kind() == reflect.Func) && rv.IsNil() {
			return err
		}
		return nil
	})
}
