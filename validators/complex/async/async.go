package complex

import (
	"sync"

	"github.com/lyonnee/hvalid"
)

// AsyncValidator 异步验证器结构体
type AsyncValidator[T any] struct {
	FieldName string // 字段名称
}

// NewAsyncValidator 创建异步验证器
func NewAsyncValidator[T any](fieldName string) *AsyncValidator[T] {
	return &AsyncValidator[T]{
		FieldName: fieldName,
	}
}

// Parallel 并行执行多个验证器
func (v *AsyncValidator[T]) Parallel(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var wg sync.WaitGroup
		errChan := make(chan error, len(validators))
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, validator := range validators {
			wg.Add(1)
			go func(validator hvalid.ValidatorFunc[T]) {
				defer wg.Done()
				if err := validator(value); err != nil {
					errChan <- err
				}
			}(validator)
		}

		// 等待所有验证器完成
		wg.Wait()
		close(errChan)

		// 收集所有错误
		for err := range errChan {
			validationErr.AddError(err.Error())
		}

		if validationErr.HasError() {
			return validationErr
		}
		return nil
	})
}

// Race 竞争执行多个验证器，返回第一个成功的结果
func (v *AsyncValidator[T]) Race(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var wg sync.WaitGroup
		successChan := make(chan struct{})
		errChan := make(chan error, len(validators))
		validationErr := hvalid.NewValidationError(v.FieldName)

		for _, validator := range validators {
			wg.Add(1)
			go func(validator hvalid.ValidatorFunc[T]) {
				defer wg.Done()
				if err := validator(value); err == nil {
					select {
					case successChan <- struct{}{}:
					default:
					}
				} else {
					errChan <- err
				}
			}(validator)
		}

		// 等待所有验证器完成或第一个成功的结果
		go func() {
			wg.Wait()
			close(successChan)
			close(errChan)
		}()

		// 检查是否有成功的验证
		if _, ok := <-successChan; ok {
			return nil
		}

		// 收集所有错误
		for err := range errChan {
			validationErr.AddError(err.Error())
		}

		return validationErr
	})
}
