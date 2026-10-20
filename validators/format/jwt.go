package format

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/lyonnee/hvalid"
)

// JWT 要求三段非空 Base64URL，且前两段是 JSON 对象。不验证签名。
func JWT[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		parts := strings.Split(string(value), ".")
		if len(parts) != 3 || !jwtObject(parts[0]) || !jwtObject(parts[1]) || !jwtBytes(parts[2]) {
			return errors.New("must be a JWT")
		}
		return nil
	}
}

func jwtObject(part string) bool {
	raw, ok := decodeBase64URL(part)
	if !ok {
		return false
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return false
	}
	_, isObject := value.(map[string]any)
	return isObject
}

func jwtBytes(part string) bool {
	_, ok := decodeBase64URL(part)
	return ok
}

func decodeBase64URL(part string) ([]byte, bool) {
	if part == "" || strings.ContainsAny(part, "+/") {
		return nil, false
	}
	enc := base64.RawURLEncoding
	if strings.ContainsRune(part, '=') {
		enc = base64.URLEncoding
	}
	raw, err := enc.DecodeString(part)
	return raw, err == nil
}
