package hvalid

import (
	"errors"
)

func Get[T any](input any, validators ...Validator[T]) (T, error) {
	return get[T](input, validators...)
}

func get[T any](input any, validators ...Validator[T]) (T, error) {
	var t T

	if res, ok := input.(T); ok {
		for _, v := range validators {
			if err := v.Validate(res); err != nil {
				return t, err
			}
		}

		return res, nil
	}

	return t, errors.New("type not match")
}
