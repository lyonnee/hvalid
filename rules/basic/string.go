package basic

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/lyonnee/hvalid"
)

func ContainsStr(substr string, errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
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

func IsIPv4(errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		if checkIPv4(field) {
			return nil
		}

		if len(errMsg) > 0 {
			return errors.New(errMsg[0])
		}

		return errors.New("the value not is ipv4 address")
	})
}

func IsIPv6(errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		if checkIPv6(field) {
			return nil
		}

		if len(errMsg) > 0 {
			return errors.New(errMsg[0])
		}

		return errors.New("the value not is ipv6 address")
	})
}

func IsUrl(errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
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

func IsEmail(errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		return Regexp(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, errMsg...)(field)
	})
}

func Regexp(pattern string, errMsg ...string) hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
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

func checkIPv4(IP string) bool {
	// 字符串这样切割
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return false
	}
	for _, s := range strs {
		if len(s) == 0 || (len(s) > 1 && s[0] == '0') {
			return false
		}
		// 直接访问字符串的值
		if s[0] < '0' || s[0] > '9' {
			return false
		}
		// 字符串转数字
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func checkIPv6(IP string) bool {
	strs := strings.Split(IP, ":")
	if len(strs) != 8 {
		return false
	}
	for _, s := range strs {
		if len(s) <= 0 || len(s) > 4 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				continue
			}
			if s[i] >= 'A' && s[i] <= 'F' {
				continue
			}
			if s[i] >= 'a' && s[i] <= 'f' {
				continue
			}
			return false
		}
	}
	return true
}
