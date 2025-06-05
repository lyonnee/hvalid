package common

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/lyonnee/hvalid/validators/primitive"
)

// PasswordValidator 密码验证器
type PasswordValidator struct {
	*primitive.StringValidator
}

// NewPasswordValidator 创建一个新的密码验证器
func NewPasswordValidator(fieldName string) *PasswordValidator {
	return &PasswordValidator{
		StringValidator: primitive.NewStringValidator(fieldName),
	}
}

// ValidateStrength 验证密码强度
func (v *PasswordValidator) ValidateStrength(minLength int) func(string) error {
	return func(s string) error {
		if len(s) < minLength {
			return fmt.Errorf("密码长度必须大于等于%d", minLength)
		}

		var (
			hasUpper   bool
			hasLower   bool
			hasNumber  bool
			hasSpecial bool
		)

		for _, char := range s {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		if !hasUpper {
			return fmt.Errorf("密码必须包含大写字母")
		}
		if !hasLower {
			return fmt.Errorf("密码必须包含小写字母")
		}
		if !hasNumber {
			return fmt.Errorf("密码必须包含数字")
		}
		if !hasSpecial {
			return fmt.Errorf("密码必须包含特殊字符")
		}

		return nil
	}
}

// ValidateComplexity 验证密码复杂度
func (v *PasswordValidator) ValidateComplexity(minTypes int) func(string) error {
	return func(s string) error {
		var types int

		if strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			types++
		}
		if strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
			types++
		}
		if strings.ContainsAny(s, "0123456789") {
			types++
		}
		if strings.ContainsAny(s, "!@#$%^&*()_+-=[]{}|;:,.<>?") {
			types++
		}

		if types < minTypes {
			return fmt.Errorf("密码必须包含至少%d种字符类型", minTypes)
		}

		return nil
	}
}

// ValidateCommon 验证是否为常见密码
func (v *PasswordValidator) ValidateCommon(commonPasswords []string) func(string) error {
	return func(s string) error {
		for _, pwd := range commonPasswords {
			if strings.EqualFold(s, pwd) {
				return fmt.Errorf("不能使用常见密码")
			}
		}
		return nil
	}
}

// ValidateRule 验证密码规则
func (v *PasswordValidator) ValidateRule(rules ...func(string) error) func(string) error {
	return func(s string) error {
		for _, rule := range rules {
			if err := rule(s); err != nil {
				return err
			}
		}
		return nil
	}
}
