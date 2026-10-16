package hvalid

import "errors"

// IsTrue 要求值为 true。
func IsTrue() ValidatorFunc[bool] {
	return func(value bool) error {
		if !value {
			return errors.New("must be true")
		}
		return nil
	}
}

// IsFalse 要求值为 false。
func IsFalse() ValidatorFunc[bool] {
	return func(value bool) error {
		if value {
			return errors.New("must be false")
		}
		return nil
	}
}
