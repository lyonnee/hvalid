package primitive

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrTextTooShort = "value length too short"
	ErrTextTooLong  = "value length too long"
)

func MinLen[T string | []byte](minLen int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l < minLen {
			return errors.New(ErrTextTooShort)
		}
		return nil
	})
}

func MaxLen[T string | []byte](maxLen int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l > maxLen {
			return errors.New(ErrTextTooLong)
		}
		return nil
	})
}
