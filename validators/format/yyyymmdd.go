package format

import (
	"errors"
	"time"

	"github.com/lyonnee/hvalid"
)

// YYYYMMDD 要求 8 位数字，且是真实公历日期。
func YYYYMMDD[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		text := string(value)
		if len(text) != 8 || !isDigits(text) {
			return errors.New("must be a YYYYMMDD date")
		}
		parsed, err := time.Parse("20060102", text)
		if err != nil || parsed.Format("20060102") != text {
			return errors.New("must be a YYYYMMDD date")
		}
		return nil
	}
}

func isDigits(text string) bool {
	for i := 0; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return false
		}
	}
	return true
}
