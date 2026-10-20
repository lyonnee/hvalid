package format

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

// MongoID 要求恰好 24 位十六进制。
func MongoID[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		text := string(value)
		if len(text) != 24 || !isHex(text) {
			return errors.New("must be a MongoDB object id")
		}
		return nil
	}
}

func isHex(text string) bool {
	for i := 0; i < len(text); i++ {
		c := text[i]
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			continue
		}
		return false
	}
	return true
}
