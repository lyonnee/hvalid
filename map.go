package hvalid

import (
	"errors"
	"fmt"
)

// MinSize 要求 map 的条目不少于 minSize。
func MinSize[K comparable, V any](minSize int) ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if len(value) < minSize {
			return fmt.Errorf("must have at least %d entries", minSize)
		}
		return nil
	}
}

// MaxSize 要求 map 的条目不超过 maxSize。
func MaxSize[K comparable, V any](maxSize int) ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if len(value) > maxSize {
			return fmt.Errorf("must have at most %d entries", maxSize)
		}
		return nil
	}
}

// MapNotEmpty 要求 map 至少有一条。
func MapNotEmpty[K comparable, V any]() ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if len(value) == 0 {
			return errors.New("must not be empty")
		}
		return nil
	}
}

// MapEmpty 要求 map 没有条目。
func MapEmpty[K comparable, V any]() ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if len(value) > 0 {
			return errors.New("must be empty")
		}
		return nil
	}
}

// HasKey 要求 map 含有 key。
func HasKey[K comparable, V any](key K) ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if _, ok := value[key]; !ok {
			return fmt.Errorf("must contain key: %v", key)
		}
		return nil
	}
}

// NoKey 要求 map 不含 key。
func NoKey[K comparable, V any](key K) ValidatorFunc[map[K]V] {
	return func(value map[K]V) error {
		if _, ok := value[key]; ok {
			return fmt.Errorf("must not contain key: %v", key)
		}
		return nil
	}
}
