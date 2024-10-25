package hvalid

import (
	"errors"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Min[T number](min T, errMsg ...string) ValidatorFunc[T] {
	return ValidatorFunc[T](func(num T) error {
		if num < min {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value is too small")
		}
		return nil
	})
}

func Max[T number](max T, errMsg ...string) ValidatorFunc[T] {
	return ValidatorFunc[T](func(num T) error {
		if num > max {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value is too big")
		}
		return nil
	})
}
