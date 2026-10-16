package hvalid

import (
	"errors"
	"fmt"
)

// MinItems 要求切片长度不少于 minLen。
func MinItems[T any](minLen int) ValidatorFunc[[]T] {
	return func(value []T) error {
		if len(value) < minLen {
			return fmt.Errorf("length must be at least %d", minLen)
		}
		return nil
	}
}

// MaxItems 要求切片长度不超过 maxLen。
func MaxItems[T any](maxLen int) ValidatorFunc[[]T] {
	return func(value []T) error {
		if len(value) > maxLen {
			return fmt.Errorf("length must be at most %d", maxLen)
		}
		return nil
	}
}

// NotEmpty 要求切片至少有一个元素。
func NotEmpty[T any]() ValidatorFunc[[]T] {
	return func(value []T) error {
		if len(value) == 0 {
			return errors.New("must not be empty")
		}
		return nil
	}
}

// Empty 要求切片没有元素。
func Empty[T any]() ValidatorFunc[[]T] {
	return func(value []T) error {
		if len(value) > 0 {
			return errors.New("must be empty")
		}
		return nil
	}
}

// Contains 要求切片含有与 element 文本形式相同的元素。
func Contains[T any](element T) ValidatorFunc[[]T] {
	return func(value []T) error {
		want := fmt.Sprintf("%v", element)
		for _, item := range value {
			if fmt.Sprintf("%v", item) == want {
				return nil
			}
		}
		return errors.New("must contain the element")
	}
}
