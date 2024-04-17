package hvalid

import (
	"errors"
	"regexp"
	"strings"
)

func ContainsStr(substr string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		ok := strings.Contains(field, substr)
		if !ok {
			return errors.New("not contains substr")
		}

		return nil
	})
}

func Email() ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		return Regexp(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, "Not email address")(field)
	})
}

func Regexp(pattern string, errMsg string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		result, _ := regexp.MatchString(pattern, field)
		if !result {
			return errors.New(errMsg)
		}
		return nil
	})
}
