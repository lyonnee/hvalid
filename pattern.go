package hvalid

import (
	"errors"
	"regexp"
	"strings"
)

// ContainsStr 要求字符串包含 sub。
func ContainsStr(sub string) ValidatorFunc[string] {
	return func(value string) error {
		if !strings.Contains(value, sub) {
			return errors.New("must contain the sub string")
		}
		return nil
	}
}

// Regexp 要求字符串匹配 pattern。
func Regexp(pattern string) ValidatorFunc[string] {
	return func(value string) error {
		matched, _ := regexp.MatchString(pattern, value)
		if !matched {
			return errors.New("must match the required pattern")
		}
		return nil
	}
}
