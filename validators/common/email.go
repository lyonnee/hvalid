package common

import (
	"fmt"
	"regexp"
	"strings"
)

// EmailValidator 邮箱验证器
type EmailValidator struct {
	FieldName string // 字段名称
}

// NewEmailValidator 创建一个新的邮箱验证器
func NewEmailValidator(fieldName string) *EmailValidator {
	return &EmailValidator{
		FieldName: fieldName,
	}
}

// Validate 验证邮箱格式
func (v *EmailValidator) Validate() func(string) error {
	return func(s string) error {
		// 移除所有空格
		s = strings.TrimSpace(s)

		// 验证基本格式
		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("无效的邮箱格式")
		}

		return nil
	}
}

// ValidateDomain 验证邮箱域名
func (v *EmailValidator) ValidateDomain(validDomains []string) func(string) error {
	return func(s string) error {
		parts := strings.Split(s, "@")
		if len(parts) != 2 {
			return fmt.Errorf("无效的邮箱格式")
		}

		domain := parts[1]
		for _, validDomain := range validDomains {
			if domain == validDomain {
				return nil
			}
		}

		return fmt.Errorf("无效的邮箱域名")
	}
}

// ValidateUsername 验证邮箱用户名
func (v *EmailValidator) ValidateUsername(maxLength int) func(string) error {
	return func(s string) error {
		parts := strings.Split(s, "@")
		if len(parts) != 2 {
			return fmt.Errorf("无效的邮箱格式")
		}

		username := parts[0]
		if len(username) > maxLength {
			return fmt.Errorf("邮箱用户名长度不能超过%d个字符", maxLength)
		}

		// 验证用户名格式
		if strings.Contains(username, "..") {
			return fmt.Errorf("邮箱用户名不能包含连续的点号")
		}

		return nil
	}
}

// ValidateDisposable 验证是否为一次性邮箱
func (v *EmailValidator) ValidateDisposable(disposableDomains []string) func(string) error {
	return func(s string) error {
		parts := strings.Split(s, "@")
		if len(parts) != 2 {
			return fmt.Errorf("无效的邮箱格式")
		}

		domain := parts[1]
		for _, disposableDomain := range disposableDomains {
			if domain == disposableDomain {
				return fmt.Errorf("不允许使用一次性邮箱")
			}
		}

		return nil
	}
}
