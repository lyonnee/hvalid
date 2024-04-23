package hvalid

import (
	"errors"
)

func Eq[T comparable](comparData T, errMsg ...string) ValidatorFunc[T] {
	return ValidatorFunc[T](func(data T) error {
		if data != comparData {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("the two values are not equal")
		}

		return nil
	})
}
