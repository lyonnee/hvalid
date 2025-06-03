package special

import (
	"fmt"
	"unicode"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrPasswordTooShort  = "length must be at least %d characters"
	ErrPasswordTooLong   = "length must be at most %d characters"
	ErrPasswordNoUpper   = "must contain uppercase letter"
	ErrPasswordNoLower   = "must contain lowercase letter"
	ErrPasswordNoNumber  = "must contain number"
	ErrPasswordNoSpecial = "must contain special character"
)

// PasswordValidator 密码验证器结构体
type PasswordValidator[T string] struct {
	MinLength      int    // 最小长度
	MaxLength      int    // 最大长度
	RequireUpper   bool   // 是否需要大写字母
	RequireLower   bool   // 是否需要小写字母
	RequireNumber  bool   // 是否需要数字
	RequireSpecial bool   // 是否需要特殊字符
	SpecialChars   []rune // 允许的特殊字符列表
	FieldName      string // 字段名称
}

// NewPasswordValidator 创建密码验证器
func NewPasswordValidator(fieldName string) *PasswordValidator[string] {
	return &PasswordValidator[string]{
		MinLength:      8,
		MaxLength:      32,
		RequireUpper:   true,
		RequireLower:   true,
		RequireNumber:  true,
		RequireSpecial: true,
		SpecialChars:   []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/'},
		FieldName:      fieldName,
	}
}

func (p *PasswordValidator[T]) Validate(password T) error {
	validationErr := hvalid.NewValidationError(p.FieldName)

	// 验证长度
	if len(password) < p.MinLength {
		validationErr.AddError(fmt.Sprintf(ErrPasswordTooShort, p.MinLength))
	}
	if len(password) > p.MaxLength {
		validationErr.AddError(fmt.Sprintf(ErrPasswordTooLong, p.MaxLength))
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	// 检查字符类型
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case p.isSpecialChar(char):
			hasSpecial = true
		}
	}

	// 验证字符要求
	if p.RequireUpper && !hasUpper {
		validationErr.AddError(ErrPasswordNoUpper)
	}
	if p.RequireLower && !hasLower {
		validationErr.AddError(ErrPasswordNoLower)
	}
	if p.RequireNumber && !hasNumber {
		validationErr.AddError(ErrPasswordNoNumber)
	}
	if p.RequireSpecial && !hasSpecial {
		validationErr.AddError(ErrPasswordNoSpecial)
	}

	if validationErr.HasError() {
		return validationErr
	}

	return nil
}

// isSpecialChar 判断字符是否为允许的特殊字符
func (p *PasswordValidator[T]) isSpecialChar(char rune) bool {
	for _, special := range p.SpecialChars {
		if char == special {
			return true
		}
	}
	return false
}

// SetMinLength 设置最小长度
func (p *PasswordValidator[T]) SetMinLength(length int) *PasswordValidator[T] {
	p.MinLength = length
	return p
}

// SetMaxLength 设置最大长度
func (p *PasswordValidator[T]) SetMaxLength(length int) *PasswordValidator[T] {
	p.MaxLength = length
	return p
}

// SetRequireUpper 设置是否需要大写字母
func (p *PasswordValidator[T]) SetRequireUpper(require bool) *PasswordValidator[T] {
	p.RequireUpper = require
	return p
}

// SetRequireLower 设置是否需要小写字母
func (p *PasswordValidator[T]) SetRequireLower(require bool) *PasswordValidator[T] {
	p.RequireLower = require
	return p
}

// SetRequireNumber 设置是否需要数字
func (p *PasswordValidator[T]) SetRequireNumber(require bool) *PasswordValidator[T] {
	p.RequireNumber = require
	return p
}

// SetRequireSpecial 设置是否需要特殊字符
func (p *PasswordValidator[T]) SetRequireSpecial(require bool) *PasswordValidator[T] {
	p.RequireSpecial = require
	return p
}

// SetSpecialChars 设置允许的特殊字符
func (p *PasswordValidator[T]) SetSpecialChars(chars []rune) *PasswordValidator[T] {
	p.SpecialChars = chars
	return p
}
