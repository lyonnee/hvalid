package format

import (
	"errors"
	"strconv"

	"github.com/lyonnee/hvalid"
)

// Port 要求整段是 0 到 65535 的十进制数，允许前导零。
func Port[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if _, err := strconv.ParseUint(string(value), 10, 16); err != nil {
			return errors.New("must be a port")
		}
		return nil
	}
}
