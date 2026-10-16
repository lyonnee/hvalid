package hvalid

import (
	"errors"
	"reflect"
)

// Eq 要求值与 want 深度相等。
func Eq[T any](want T) ValidatorFunc[T] {
	return func(got T) error {
		if !reflect.DeepEqual(got, want) {
			return errors.New("the two values are not equal")
		}
		return nil
	}
}
