package primitive

import (
	"fmt"
	"time"

	"github.com/lyonnee/hvalid"
)

// 预定义错误信息
const (
	ErrTimeBefore = "must be before %v"
	ErrTimeAfter  = "must be after %v"
	ErrTimeEqual  = "must be equal to %v"
)

// TimeValidator 时间验证器结构体
type TimeValidator struct {
	FieldName string // 字段名称
}

// NewTimeValidator 创建时间验证器
func NewTimeValidator(fieldName string) *TimeValidator {
	return &TimeValidator{
		FieldName: fieldName,
	}
}

// Before 验证是否在指定时间之前
func (v *TimeValidator) Before(t time.Time) hvalid.ValidatorFunc[time.Time] {
	return hvalid.ValidatorFunc[time.Time](func(value time.Time) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !value.Before(t) {
			validationErr.AddError(fmt.Sprintf(ErrTimeBefore, t))
			return validationErr
		}
		return nil
	})
}

// After 验证是否在指定时间之后
func (v *TimeValidator) After(t time.Time) hvalid.ValidatorFunc[time.Time] {
	return hvalid.ValidatorFunc[time.Time](func(value time.Time) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !value.After(t) {
			validationErr.AddError(fmt.Sprintf(ErrTimeAfter, t))
			return validationErr
		}
		return nil
	})
}

// Equal 验证是否等于指定时间
func (v *TimeValidator) Equal(t time.Time) hvalid.ValidatorFunc[time.Time] {
	return hvalid.ValidatorFunc[time.Time](func(value time.Time) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if !value.Equal(t) {
			validationErr.AddError(fmt.Sprintf(ErrTimeEqual, t))
			return validationErr
		}
		return nil
	})
}

// Between 验证是否在指定时间范围内
func (v *TimeValidator) Between(start, end time.Time) hvalid.ValidatorFunc[time.Time] {
	return hvalid.ValidatorFunc[time.Time](func(value time.Time) error {
		validationErr := hvalid.NewValidationError(v.FieldName)

		if value.Before(start) || value.After(end) {
			validationErr.AddError(fmt.Sprintf("must be between %v and %v", start, end))
			return validationErr
		}
		return nil
	})
}
