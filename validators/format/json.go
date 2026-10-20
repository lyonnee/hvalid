package format

import (
	"encoding/json"
	"errors"

	"github.com/lyonnee/hvalid"
)

// JSON 要求非空且是 JSON。
func JSON[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if value == "" || !json.Valid([]byte(value)) {
			return errors.New("must be JSON")
		}
		return nil
	}
}
