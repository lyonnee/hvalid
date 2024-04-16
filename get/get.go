package get

import (
	"errors"

	"github.com/lyonnee/hval"
)

func Get[T any](input any, validators ...hval.Validator[T]) (T, error) {
	return get[T](input, validators...)
}

func get[T any](input any, validators ...hval.Validator[T]) (T, error) {
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
