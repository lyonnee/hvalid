package special

import (
	"errors"
	"unicode"
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
}

// NewPasswordValidator 创建密码验证器
func NewPasswordValidator(
	minLen int,
	maxLen int,
	reqUpper bool,
	reqLower bool,
	reqNumber bool,
	reqSpecial bool,
	specialChars []rune, // []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/'}
) *PasswordValidator[string] {
	return &PasswordValidator[string]{
		MinLength:      minLen,
		MaxLength:      maxLen,
		RequireUpper:   reqUpper,
		RequireLower:   reqLower,
		RequireNumber:  reqNumber,
		RequireSpecial: reqSpecial,
		SpecialChars:   specialChars,
	}
}

func (p *PasswordValidator[T]) Validate(password T) error {
	if len(password) < p.MinLength {
		return errors.New("密码长度不能小于" + string(p.MinLength))
	}
	if len(password) > p.MaxLength {
		return errors.New("密码长度不能大于" + string(p.MaxLength))
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

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

	if p.RequireUpper && !hasUpper {
		return errors.New("密码必须包含大写字母")
	}
	if p.RequireLower && !hasLower {
		return errors.New("密码必须包含小写字母")
	}
	if p.RequireNumber && !hasNumber {
		return errors.New("密码必须包含数字")
	}
	if p.RequireSpecial && !hasSpecial {
		return errors.New("密码必须包含特殊字符")
	}

	return errors.New("密码格式正确")
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
