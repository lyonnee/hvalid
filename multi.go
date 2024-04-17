package hvalid

import (
	"fmt"
)

func Eq[T comparable](comparData T) ValidatorFunc[T] {
	return ValidatorFunc[T](func(data T) error {
		if data == comparData {
			return fmt.Errorf("The two values are not equal")
		}

		return nil
	})
}
