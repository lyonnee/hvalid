package hvalid

import (
	"errors"
	"fmt"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Min 要求数值大于等于 min。
func Min[T number](min T) ValidatorFunc[T] {
	return func(value T) error {
		if value < min {
			return fmt.Errorf("must be greater than or equal to %v", min)
		}
		return nil
	}
}

// Max 要求数值小于等于 max。
func Max[T number](max T) ValidatorFunc[T] {
	return func(value T) error {
		if value > max {
			return fmt.Errorf("must be less than or equal to %v", max)
		}
		return nil
	}
}

// Range 要求数值落在 min 与 max 之间，含两端。
func Range[T number](min, max T) ValidatorFunc[T] {
	return func(value T) error {
		if value < min || value > max {
			return fmt.Errorf("must be between %v and %v", min, max)
		}
		return nil
	}
}

// Positive 要求数值大于 0。
func Positive[T number]() ValidatorFunc[T] {
	return func(value T) error {
		if value <= 0 {
			return errors.New("must be positive")
		}
		return nil
	}
}

// Negative 要求数值小于 0。
func Negative[T number]() ValidatorFunc[T] {
	return func(value T) error {
		if value >= 0 {
			return errors.New("must be negative")
		}
		return nil
	}
}
