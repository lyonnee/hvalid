package format

import "github.com/lyonnee/hvalid"

// Longitude 要求十进制经度，范围 -180 到 180。
func Longitude[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		return decimalRange(string(value), -180, 180, "must be a longitude")
	}
}
