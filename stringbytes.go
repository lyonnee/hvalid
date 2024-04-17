package hvalid

import "fmt"

func MinLen[T string | []byte](minLen int) ValidatorFunc[T] {
	return ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l < minLen {
			return fmt.Errorf("The length of the input data is too short, data len: %d, minLen: %d", l, minLen)
		}
		return nil
	})
}

func MaxLen[T string | []byte](maxLen int) ValidatorFunc[T] {
	return ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l > maxLen {
			return fmt.Errorf("The length of the input data is too long, data len: %d, maxLen: %d", l, maxLen)
		}
		return nil
	})
}
