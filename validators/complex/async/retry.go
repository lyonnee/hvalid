package complex

import (
	"fmt"
	"time"

	"github.com/lyonnee/hvalid"
)

// RetryValidator 重试验证器结构体
type RetryValidator[T any] struct {
	FieldName string // 字段名称
}

// NewRetryValidator 创建重试验证器
func NewRetryValidator[T any](fieldName string) *RetryValidator[T] {
	return &RetryValidator[T]{
		FieldName: fieldName,
	}
}

// WithRetry 使用重试机制执行验证
func (v *RetryValidator[T]) WithRetry(validator hvalid.ValidatorFunc[T], maxRetries int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var lastErr error
		for i := 0; i <= maxRetries; i++ {
			if err := validator(value); err == nil {
				return nil
			} else {
				lastErr = err
			}
		}
		return fmt.Errorf("validation failed after %d retries: %v", maxRetries, lastErr)
	})
}

// WithBackoff 使用指数退避重试机制执行验证
func (v *RetryValidator[T]) WithBackoff(validator hvalid.ValidatorFunc[T], maxRetries int, initialDelay time.Duration) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var lastErr error
		delay := initialDelay

		for i := 0; i <= maxRetries; i++ {
			if err := validator(value); err == nil {
				return nil
			} else {
				lastErr = err
				if i < maxRetries {
					time.Sleep(delay)
					delay *= 2 // 指数退避
				}
			}
		}
		return fmt.Errorf("validation failed after %d retries with backoff: %v", maxRetries, lastErr)
	})
}

// WithCondition 根据条件决定是否重试
func (v *RetryValidator[T]) WithCondition(validator hvalid.ValidatorFunc[T], shouldRetry func(error) bool, maxRetries int) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var lastErr error
		for i := 0; i <= maxRetries; i++ {
			if err := validator(value); err == nil {
				return nil
			} else {
				lastErr = err
				if !shouldRetry(err) {
					return err
				}
			}
		}
		return fmt.Errorf("validation failed after %d retries: %v", maxRetries, lastErr)
	})
}
