package format

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

// ISBN 要求去掉连字符和空格后，是校验位正确的 ISBN-10 或 ISBN-13。
func ISBN[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if !isISBN(string(value)) {
			return errors.New("must be an ISBN")
		}
		return nil
	}
}

func isISBN(text string) bool {
	compact := make([]rune, 0, len(text))
	for _, r := range text {
		if r == '-' || r == ' ' {
			continue
		}
		compact = append(compact, r)
	}
	switch len(compact) {
	case 10:
		return isbn10(compact)
	case 13:
		return isbn13(compact)
	default:
		return false
	}
}

func isbn10(rs []rune) bool {
	sum := 0
	for i, r := range rs {
		var digit int
		switch {
		case r >= '0' && r <= '9':
			digit = int(r - '0')
		case i == 9 && (r == 'X' || r == 'x'):
			digit = 10
		default:
			return false
		}
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}

func isbn13(rs []rune) bool {
	prefix := string(rs[:3])
	if prefix != "978" && prefix != "979" {
		return false
	}
	sum := 0
	for i, r := range rs {
		if r < '0' || r > '9' {
			return false
		}
		digit := int(r - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	return sum%10 == 0
}
