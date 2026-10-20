package format

import (
	"errors"
	"strings"

	"github.com/lyonnee/hvalid"
)

const crockford = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// ULID 要求 26 位 Crockford Base32，拒绝 I、L、O、U。
func ULID[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		text := string(value)
		if len(text) != 26 || !isCrockford(text) {
			return errors.New("must be a ULID")
		}
		return nil
	}
}

func isCrockford(text string) bool {
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c >= 'a' && c <= 'z' {
			c = c - 'a' + 'A'
		}
		if !strings.ContainsRune(crockford, rune(c)) {
			return false
		}
	}
	return true
}
