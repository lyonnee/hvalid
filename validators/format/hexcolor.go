package format

import (
	"errors"
	"regexp"

	"github.com/lyonnee/hvalid"
)

var hexColorPattern = regexp.MustCompile(`(?i)^#(?:[0-9a-f]{3}|[0-9a-f]{6}|[0-9a-f]{8})$`)

// HexColor 要求 #RGB、#RRGGBB 或 #RRGGBBAA。
func HexColor[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if !hexColorPattern.MatchString(string(value)) {
			return errors.New("must be a hex color")
		}
		return nil
	}
}
