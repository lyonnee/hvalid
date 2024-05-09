package hvalid

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

func ContainsStr(substr string, errMsg ...string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		ok := strings.Contains(field, substr)
		if !ok {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("not contain the sub string")
		}

		return nil
	})
}

func IsUrl(errMsg ...string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		err := errors.New("the value not is url")
		if len(errMsg) > 0 {
			err = errors.New(errMsg[0])
		}

		_, parseErr := url.ParseRequestURI(field)
		if parseErr != nil {
			return err
		}

		u, parseErr := url.Parse(field)
		if parseErr != nil {
			return err
		}

		if u.Scheme == "" && u.Host == "" {
			return err
		}

		return nil
	})
}

func IsEmail(errMsg ...string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		return Regexp(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, errMsg...)(field)
	})
}

func Regexp(pattern string, errMsg ...string) ValidatorFunc[string] {
	return ValidatorFunc[string](func(field string) error {
		result, _ := regexp.MatchString(pattern, field)
		if !result {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("not an email address")
		}
		return nil
	})
}
