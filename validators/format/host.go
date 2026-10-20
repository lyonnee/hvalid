package format

import (
	"errors"
	"net"

	"github.com/lyonnee/hvalid"
)

// Host 要求是 IP 或 DNS 名。不接受带方括号的 IPv6。
func Host[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if net.ParseIP(string(value)) != nil {
			return nil
		}
		if DNSName[T]()(value) == nil {
			return nil
		}
		return errors.New("must be a host")
	}
}
