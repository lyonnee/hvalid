package format

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/lyonnee/hvalid"
)

// Base64 要求标准 Base64，填充正确，且不含空白。
func Base64[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		text := string(value)
		if text == "" || strings.ContainsAny(text, " \t\r\n") {
			return errors.New("must be base64")
		}
		if _, err := base64.StdEncoding.DecodeString(text); err != nil {
			return errors.New("must be base64")
		}
		return nil
	}
}
