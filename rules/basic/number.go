package basic

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

func Min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](min T, errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(num T) error {
		if num < min {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value is too small")
		}
		return nil
	})
}

func Max[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](max T, errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(num T) error {
		if num > max {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value is too big")
		}
		return nil
	})
}
