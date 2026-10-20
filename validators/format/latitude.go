package format

import "github.com/lyonnee/hvalid"

// Latitude 要求十进制纬度，范围 -90 到 90。
func Latitude[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		return decimalRange(string(value), -90, 90, "must be a latitude")
	}
}
