package basic

import (
	"errors"

	"github.com/lyonnee/hvalid"
)

func MinLen[T string | []byte](minLen int, errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l < minLen {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value length too short")
		}
		return nil
	})
}

func MaxLen[T string | []byte](maxLen int, errMsg ...string) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(filed T) error {
		l := len(filed)
		if l > maxLen {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("value length too long")
		}
		return nil
	})
}
