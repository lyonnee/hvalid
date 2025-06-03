package primitive

import (
	"bytes"
	"errors"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrBytesNotContains = "must contain the sub byte slice"
)

func ContainsBytes(subslice []byte) hvalid.ValidatorFunc[[]byte] {
	return hvalid.ValidatorFunc[[]byte](func(field []byte) error {
		ok := bytes.Contains(field, subslice)
		if !ok {
			return errors.New(ErrBytesNotContains)
		}

		return nil
	})
}
