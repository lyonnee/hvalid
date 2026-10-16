package hvalid

import "errors"

// MinLen 要求文本长度不少于 minLen。
func MinLen[T ~string | ~[]byte](minLen int) ValidatorFunc[T] {
	return func(value T) error {
		if len(value) < minLen {
			return errors.New("value length too short")
		}
		return nil
	}
}

// MaxLen 要求文本长度不超过 maxLen。
func MaxLen[T ~string | ~[]byte](maxLen int) ValidatorFunc[T] {
	return func(value T) error {
		if len(value) > maxLen {
			return errors.New("value length too long")
		}
		return nil
	}
}
