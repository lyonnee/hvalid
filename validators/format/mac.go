package format

import (
	"errors"
	"net"

	"github.com/lyonnee/hvalid"
)

// MAC 要求 net.ParseMAC 能够解析。
func MAC[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if _, err := net.ParseMAC(string(value)); err != nil {
			return errors.New("must be a MAC address")
		}
		return nil
	}
}
