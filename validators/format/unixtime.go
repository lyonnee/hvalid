package format

import (
	"errors"
	"strconv"

	"github.com/lyonnee/hvalid"
)

// UnixTime 要求可选负号加数字，且能放进 int64。
func UnixTime[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		text := string(value)
		if !isUnixTime(text) {
			return errors.New("must be a unix time")
		}
		return nil
	}
}

func isUnixTime(text string) bool {
	if text == "" {
		return false
	}
	i := 0
	if text[0] == '-' {
		i++
	}
	if i == len(text) {
		return false
	}
	for ; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return false
		}
	}
	_, err := strconv.ParseInt(text, 10, 64)
	return err == nil
}
