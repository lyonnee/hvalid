package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CreditCardValidator 信用卡号验证器
type CreditCardValidator struct {
	FieldName string // 字段名称
}

// NewCreditCardValidator 创建一个新的信用卡号验证器
func NewCreditCardValidator(fieldName string) *CreditCardValidator {
	return &CreditCardValidator{
		FieldName: fieldName,
	}
}

// Validate 验证信用卡号格式
func (v *CreditCardValidator) Validate() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 验证长度（13-19位）
		if len(s) < 13 || len(s) > 19 {
			return fmt.Errorf("信用卡号长度必须在13-19位之间")
		}

		// 验证是否都是数字
		for _, c := range s {
			if c < '0' || c > '9' {
				return fmt.Errorf("信用卡号只能包含数字")
			}
		}

		return nil
	}
}

// ValidateCardType 验证信用卡号卡组织
func (v *CreditCardValidator) ValidateCardType() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// 定义卡组织规则
		cardRules := map[string]struct {
			pattern string
			name    string
		}{
			"visa": {
				pattern: "^4[0-9]{12}(?:[0-9]{3})?$",
				name:    "Visa",
			},
			"mastercard": {
				pattern: "^5[1-5][0-9]{14}$",
				name:    "MasterCard",
			},
			"amex": {
				pattern: "^3[47][0-9]{13}$",
				name:    "American Express",
			},
			"discover": {
				pattern: "^6(?:011|5[0-9]{2})[0-9]{12}$",
				name:    "Discover",
			},
		}

		// 验证卡组织
		for _, rule := range cardRules {
			matched, err := regexp.MatchString(rule.pattern, s)
			if err != nil {
				return fmt.Errorf("验证失败: %v", err)
			}
			if matched {
				return nil
			}
		}

		return fmt.Errorf("无效的卡组织")
	}
}

// ValidateLuhn 验证信用卡号Luhn算法
func (v *CreditCardValidator) ValidateLuhn() func(string) error {
	return func(s string) error {
		// 移除所有空格和连字符
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "-", "")

		// Luhn算法验证
		sum := 0
		alternate := false

		// 从右向左遍历
		for i := len(s) - 1; i >= 0; i-- {
			digit, _ := strconv.Atoi(string(s[i]))

			if alternate {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}

			sum += digit
			alternate = !alternate
		}

		if sum%10 != 0 {
			return fmt.Errorf("Luhn算法验证失败")
		}

		return nil
	}
}

// ValidateExpiryDate 验证信用卡号有效期
func (v *CreditCardValidator) ValidateExpiryDate() func(string) error {
	return func(s string) error {
		// 验证格式（MM/YY）
		pattern := `^(0[1-9]|1[0-2])/([0-9]{2})$`
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return fmt.Errorf("验证失败: %v", err)
		}
		if !matched {
			return fmt.Errorf("无效的有效期格式")
		}

		// 验证是否过期
		parts := strings.Split(s, "/")
		month, _ := strconv.Atoi(parts[0])
		year, _ := strconv.Atoi("20" + parts[1])

		// 这里只是示例，实际应该使用当前时间进行比较
		if year < 2024 || (year == 2024 && month < 1) {
			return fmt.Errorf("信用卡已过期")
		}

		return nil
	}
}
