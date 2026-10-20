package format

import (
	"errors"
	"regexp"

	"github.com/lyonnee/hvalid"
)

var uuidPattern = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-8][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

// UUID 要求值是 RFC 4122 UUID。versions 为空时接受版本 1 到 8，否则版本半字节必须命中其中之一。
func UUID[T ~string](versions ...int) hvalid.ValidatorFunc[T] {
	want := append([]int(nil), versions...)
	return func(value T) error {
		text := string(value)
		if !uuidPattern.MatchString(text) || !uuidVersion(text, want) {
			return errors.New("must be a UUID")
		}
		return nil
	}
}

func uuidVersion(text string, versions []int) bool {
	if len(versions) == 0 {
		return true
	}
	version := int(text[14] - '0')
	for _, item := range versions {
		if item == version {
			return true
		}
	}
	return false
}
