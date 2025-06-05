package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lyonnee/hvalid/validators/primitive"
)

// PostcodeValidator 邮政编码验证器
type PostcodeValidator struct {
	*primitive.StringValidator
}

// NewPostcodeValidator 创建一个新的邮政编码验证器
func NewPostcodeValidator(fieldName string) *PostcodeValidator {
	return &PostcodeValidator{
		StringValidator: primitive.NewStringValidator(fieldName),
	}
}

// ValidateCN 验证中国邮政编码
func (v *PostcodeValidator) ValidateCN() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证长度（6位）
		if len(s) != 6 {
			return fmt.Errorf("中国邮政编码必须是6位数字")
		}

		// 验证是否都是数字
		for _, c := range s {
			if c < '0' || c > '9' {
				return fmt.Errorf("中国邮政编码只能包含数字")
			}
		}

		// 验证第一位数字（1-9）
		if s[0] < '1' || s[0] > '9' {
			return fmt.Errorf("中国邮政编码第一位必须是1-9")
		}

		return nil
	}
}

// ValidateUS 验证美国邮政编码
func (v *PostcodeValidator) ValidateUS() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证格式（5位数字或5位数字-4位数字）
		pattern := `^\d{5}(-\d{4})?$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("美国邮政编码格式无效，应为5位数字或5位数字-4位数字")
		}

		return nil
	}
}

// ValidateUK 验证英国邮政编码
func (v *PostcodeValidator) ValidateUK() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证格式（1-2个字母+1-2个数字+1个数字+2个字母）
		pattern := `^[A-Z]{1,2}\d[A-Z\d]? ?\d[A-Z]{2}$`
		matched, err := regexp.MatchString(pattern, strings.ToUpper(s))
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("英国邮政编码格式无效")
		}

		return nil
	}
}

// ValidateCA 验证加拿大邮政编码
func (v *PostcodeValidator) ValidateCA() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证格式（字母数字字母 数字字母数字）
		pattern := `^[A-Z]\d[A-Z]\d[A-Z]\d$`
		matched, err := regexp.MatchString(pattern, strings.ToUpper(s))
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("加拿大邮政编码格式无效，应为字母数字字母数字字母数字")
		}

		return nil
	}
}

// ValidateAU 验证澳大利亚邮政编码
func (v *PostcodeValidator) ValidateAU() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证格式（4位数字）
		pattern := `^\d{4}$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("澳大利亚邮政编码必须是4位数字")
		}

		return nil
	}
}

// ValidateJP 验证日本邮政编码
func (v *PostcodeValidator) ValidateJP() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 验证格式（7位数字）
		pattern := `^\d{7}$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("日本邮政编码必须是7位数字")
		}

		return nil
	}
}

// ValidateFormat 验证自定义格式的邮政编码
func (v *PostcodeValidator) ValidateFormat(pattern string) func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.ReplaceAll(s, " ", "")

		// 验证格式
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("邮政编码格式无效")
		}

		return nil
	}
}
