package hvalid

import (
	"errors"
	"unicode/utf8"
)

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

// MinRunes 要求符文数不少于 n。
func MinRunes[T ~string](n int) ValidatorFunc[T] {
	return func(value T) error {
		if utf8.RuneCountInString(string(value)) < n {
			return errors.New("value rune length too short")
		}
		return nil
	}
}

// MaxRunes 要求符文数不超过 n。
func MaxRunes[T ~string](n int) ValidatorFunc[T] {
	return func(value T) error {
		if utf8.RuneCountInString(string(value)) > n {
			return errors.New("value rune length too long")
		}
		return nil
	}
}
