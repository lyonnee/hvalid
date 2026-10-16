package primitive

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrNotIPv4  = "must be a valid IPv4 address"
	ErrNotIPv6  = "must be a valid IPv6 address"
	ErrNotURL   = "must be a valid URL"
	ErrNotEmail = "must be a valid email address"
)

// StringValidator 字符串验证器结构体
type StringValidator struct {
	FieldName string // 字段名称
}

// NewStringValidator 创建字符串验证器
func NewStringValidator(fieldName string) *StringValidator {
	return &StringValidator{
		FieldName: fieldName,
	}
}

// IsIPv4 验证是否为IPv4地址
func (v *StringValidator) IsIPv4() hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !checkIPv4(field) {
			validationErr.AddError(ErrNotIPv4)
			return validationErr
		}
		return nil
	})
}

// IsIPv6 验证是否为IPv6地址
func (v *StringValidator) IsIPv6() hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !checkIPv6(field) {
			validationErr.AddError(ErrNotIPv6)
			return validationErr
		}
		return nil
	})
}

// IsURL 验证是否为URL
func (v *StringValidator) IsURL() hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		_, parseErr := url.ParseRequestURI(field)
		if parseErr != nil {
			validationErr.AddError(ErrNotURL)
			return validationErr
		}

		u, parseErr := url.Parse(field)
		if parseErr != nil {
			validationErr.AddError(ErrNotURL)
			return validationErr
		}

		if u.Scheme == "" && u.Host == "" {
			validationErr.AddError(ErrNotURL)
			return validationErr
		}

		return nil
	})
}

// IsEmail 验证是否为邮箱地址
func (v *StringValidator) IsEmail() hvalid.ValidatorFunc[string] {
	return hvalid.ValidatorFunc[string](func(field string) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, field)
		if !result {
			validationErr.AddError(ErrNotEmail)
			return validationErr
		}
		return nil
	})
}

// checkIPv4 检查是否为有效的IPv4地址
func checkIPv4(IP string) bool {
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return false
	}
	for _, s := range strs {
		if len(s) == 0 || (len(s) > 1 && s[0] == '0') {
			return false
		}
		if s[0] < '0' || s[0] > '9' {
			return false
		}
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

// checkIPv6 检查是否为有效的IPv6地址
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
