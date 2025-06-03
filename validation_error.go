package hvalid

import (
	"fmt"
	"strings"
)

// GetDefaultErrorMsg 获取默认错误信息
func GetDefaultErrorMsg(fieldName string, rule string) string {
	return fmt.Sprintf("%s %s", fieldName, rule)
}

// ValidationError 表示验证错误
type ValidationError struct {
	Field  string   // 字段名称
	Errors []string // 错误信息列表
}

// Error 实现 error 接口
func (e *ValidationError) Error() string {
	if len(e.Errors) == 0 {
		return ""
	}
	return fmt.Sprintf("%s: %s", e.Field, strings.Join(e.Errors, "; "))
}

// NewValidationError 创建新的验证错误
func NewValidationError(field string) *ValidationError {
	return &ValidationError{
		Field:  field,
		Errors: make([]string, 0),
	}
}

// AddError 添加错误信息
func (e *ValidationError) AddError(err string) {
	e.Errors = append(e.Errors, err)
}

// HasError 检查是否有错误
func (e *ValidationError) HasError() bool {
	return len(e.Errors) > 0
}
