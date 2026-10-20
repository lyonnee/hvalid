package format

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/lyonnee/hvalid"
)

var rgbPattern = regexp.MustCompile(`^rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)$`)

// RGB 要求 rgb(r, g, b)，三个分量都是 0 到 255 的整数。
func RGB[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		parts := rgbPattern.FindStringSubmatch(string(value))
		if parts == nil {
			return errors.New("must be an RGB color")
		}
		for _, part := range parts[1:] {
			n, err := strconv.Atoi(part)
			if err != nil || n > 255 {
				return errors.New("must be an RGB color")
			}
		}
		return nil
	}
}
