package format

import (
	"errors"
	"time"

	"github.com/lyonnee/hvalid"
)

// RFC3339 要求 RFC3339 时间，可以带小数秒。
func RFC3339[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if _, err := time.Parse(time.RFC3339, string(value)); err != nil {
			return errors.New("must be RFC3339")
		}
		return nil
	}
}
