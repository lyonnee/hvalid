package format

import (
	"errors"
	"strings"

	"github.com/lyonnee/hvalid"
)

// DNSName 要求 RFC 1123 主机名。允许单标签，不接受结尾点和下划线。
func DNSName[T ~string]() hvalid.ValidatorFunc[T] {
	return func(value T) error {
		if !isDNSName(string(value)) {
			return errors.New("must be a DNS name")
		}
		return nil
	}
}

func isDNSName(text string) bool {
	if text == "" || len(text) > 253 || strings.HasSuffix(text, ".") {
		return false
	}
	for _, label := range strings.Split(text, ".") {
		if !isDNSLabel(label) {
			return false
		}
	}
	return true
}

func isDNSLabel(label string) bool {
	n := len(label)
	if n < 1 || n > 63 {
		return false
	}
	if label[0] == '-' || label[n-1] == '-' {
		return false
	}
	for i := 0; i < n; i++ {
		c := label[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' {
			continue
		}
		return false
	}
	return true
}
