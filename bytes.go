package hvalid

import (
	"bytes"
	"errors"
)

func ContainsBytes(subslice []byte) ValidatorFunc[[]byte] {
	return ValidatorFunc[[]byte](func(field []byte) error {
		ok := bytes.Contains(field, subslice)
		if !ok {
			return errors.New("not contains sub byte slice")
		}

		return nil
	})
}
