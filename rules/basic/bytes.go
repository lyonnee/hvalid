package basic

import (
	"bytes"
	"errors"

	"github.com/lyonnee/hvalid"
)

func ContainsBytes(subslice []byte) hvalid.ValidatorFunc[[]byte] {
	return hvalid.ValidatorFunc[[]byte](func(field []byte) error {
		ok := bytes.Contains(field, subslice)
		if !ok {
			return errors.New("not contains sub byte slice")
		}

		return nil
	})
}
