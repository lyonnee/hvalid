package hvalid

import (
	"fmt"
	"strings"
)

// OneOf 要求值与允许列表中的某一项精确相等，区分大小写。
func OneOf[T ~string](allowed ...T) ValidatorFunc[T] {
	return func(value T) error {
		for _, item := range allowed {
			if value == item {
				return nil
			}
		}
		parts := make([]string, len(allowed))
		for i, item := range allowed {
			parts[i] = string(item)
		}
		return fmt.Errorf("must be one of: %s", strings.Join(parts, ", "))
	}
}
