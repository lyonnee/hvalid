package format

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/lyonnee/hvalid"
)

// DataURI 要求 data: 前缀。带 ;base64 时载荷必须是标准 Base64。
func DataURI[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if !isDataURI(string(value)) {
			return errors.New("must be a data URI")
		}
		return nil
	}
}

func isDataURI(text string) bool {
	if len(text) < 5 || !strings.EqualFold(text[:5], "data:") {
		return false
	}
	rest := text[5:]
	comma := strings.IndexByte(rest, ',')
	if comma < 0 {
		return false
	}
	meta := rest[:comma]
	payload := rest[comma+1:]
	if strings.HasSuffix(strings.ToLower(meta), ";base64") {
		if _, err := base64.StdEncoding.DecodeString(payload); err != nil {
			return false
		}
	}
	return true
}
