package complex

import (
	"context"
	"fmt"
	"time"

	"github.com/lyonnee/hvalid"
)

// TimeoutValidator 超时验证器结构体
type TimeoutValidator[T any] struct {
	FieldName string // 字段名称
}

// NewTimeoutValidator 创建超时验证器
func NewTimeoutValidator[T any](fieldName string) *TimeoutValidator[T] {
	return &TimeoutValidator[T]{
		FieldName: fieldName,
	}
}

// WithTimeout 使用超时机制执行验证
func (v *TimeoutValidator[T]) WithTimeout(validator hvalid.ValidatorFunc[T], timeout time.Duration) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		errChan := make(chan error, 1)
		go func() {
			errChan <- validator(value)
		}()

		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return fmt.Errorf("validation timed out after %v", timeout)
		}
	})
}

// WithDeadline 使用截止时间执行验证
func (v *TimeoutValidator[T]) WithDeadline(validator hvalid.ValidatorFunc[T], deadline time.Time) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		errChan := make(chan error, 1)
		go func() {
			errChan <- validator(value)
		}()

		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return fmt.Errorf("validation deadline exceeded at %v", deadline)
		}
	})
}

// WithContext 使用上下文执行验证
func (v *TimeoutValidator[T]) WithContext(validator hvalid.ValidatorFunc[T], ctx context.Context) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		errChan := make(chan error, 1)
		go func() {
			errChan <- validator(value)
		}()

		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return fmt.Errorf("validation cancelled: %v", ctx.Err())
		}
	})
}
