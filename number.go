package hvalid

import (
	"fmt"
)

func Min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](min T) ValidatorFunc[T] {
	return ValidatorFunc[T](func(num T) error {
		if num < min {
			return fmt.Errorf("The num is less than the minimum value, mininum value: %v, input num: %v", min, num)
		}
		return nil
	})
}

func Max[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](max T) ValidatorFunc[T] {
	return ValidatorFunc[T](func(num T) error {
		if num > max {
			return fmt.Errorf("The num is greater than the maximum value, maximum value: %v, input num: %v", max, num)
		}
		return nil
	})
}
