package common

import (
	"fmt"
	"regexp"
	"strings"
)

// PhoneValidator 手机号验证器
type PhoneValidator struct {
	FieldName string // 字段名称
}

// NewPhoneValidator 创建一个新的手机号验证器
func NewPhoneValidator(fieldName string) *PhoneValidator {
	return &PhoneValidator{
		FieldName: fieldName,
	}
}

// ValidateCN 验证中国手机号格式
func (v *PhoneValidator) ValidateCN() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 验证长度和格式
		if len(s) != 11 {
			return fmt.Errorf("手机号长度必须为11位")
		}

		// 验证是否以1开头
		if !strings.HasPrefix(s, "1") {
			return fmt.Errorf("手机号必须以1开头")
		}

		// 验证第二位是否为3-9
		if s[1] < '3' || s[1] > '9' {
			return fmt.Errorf("手机号第二位必须是3-9之间的数字")
		}

		// 验证剩余位是否都是数字
		for i := 2; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return fmt.Errorf("手机号只能包含数字")
			}
		}

		return nil
	}
}

// ValidateInternational 验证国际手机号格式
func (v *PhoneValidator) ValidateInternational() func(string) error {
	return func(s string) error {
		// 验证国际格式
		pattern := `^\+[1-9]\d{0,3}-[1-9]\d{0,3}-\d{3,4}-\d{4}$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("无效的国际手机号格式")
		}
		return nil
	}
}

// ValidateOperator 验证手机号运营商
func (v *PhoneValidator) ValidateOperator(operators map[string]string) func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 验证运营商号段
		prefix := s[:3]
		if _, ok := operators[prefix]; !ok {
			return fmt.Errorf("无效的运营商号段")
		}

		return nil
	}
}

// ValidateFormat 验证手机号格式化和清理
func (v *PhoneValidator) ValidateFormat() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 验证清理后的格式
		if len(s) != 11 {
			return fmt.Errorf("手机号长度必须为11位")
		}

		// 验证是否都是数字
		for _, c := range s {
			if c < '0' || c > '9' {
				return fmt.Errorf("手机号只能包含数字")
			}
		}

		return nil
	}
}
