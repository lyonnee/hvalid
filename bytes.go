package hvalid

import (
	"bytes"
	"errors"
)

// ContainsBytes 要求字节切片包含 sub。
func ContainsBytes(sub []byte) ValidatorFunc[[]byte] {
	return func(value []byte) error {
		if !bytes.Contains(value, sub) {
			return errors.New("must contain the sub byte slice")
		}
		return nil
	}
}
