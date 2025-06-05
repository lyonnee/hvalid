package complex

import (
	"fmt"
	"sync"

	"github.com/lyonnee/hvalid"
)

// BatchValidator 批量验证器结构体
type BatchValidator[T any] struct {
	FieldName string // 字段名称
}

// NewBatchValidator 创建批量验证器
func NewBatchValidator[T any](fieldName string) *BatchValidator[T] {
	return &BatchValidator[T]{
		FieldName: fieldName,
	}
}

// ValidateAll 验证所有值
func (v *BatchValidator[T]) ValidateAll(values []T, validator hvalid.ValidatorFunc[T]) error {
	validationErr := hvalid.NewValidationError(v.FieldName)

	for i, value := range values {
		if err := validator(value); err != nil {
			validationErr.AddError(fmt.Sprintf("element[%d]: %v", i, err))
		}
	}

	if validationErr.HasError() {
		return validationErr
	}
	return nil
}

// ValidateAllParallel 并行验证所有值
func (v *BatchValidator[T]) ValidateAllParallel(values []T, validator hvalid.ValidatorFunc[T]) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(values))
	validationErr := hvalid.NewValidationError(v.FieldName)

	for i, value := range values {
		wg.Add(1)
		go func(index int, val T) {
			defer wg.Done()
			if err := validator(val); err != nil {
				errChan <- fmt.Errorf("element[%d]: %v", index, err)
			}
		}(i, value)
	}

	// 等待所有验证完成
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
}

// ValidateAny 验证任意一个值
func (v *BatchValidator[T]) ValidateAny(values []T, validator hvalid.ValidatorFunc[T]) error {
	validationErr := hvalid.NewValidationError(v.FieldName)

	for i, value := range values {
		if err := validator(value); err == nil {
			return nil
		} else {
			validationErr.AddError(fmt.Sprintf("element[%d]: %v", i, err))
		}
	}

	return validationErr
}

// ValidateAnyParallel 并行验证任意一个值
func (v *BatchValidator[T]) ValidateAnyParallel(values []T, validator hvalid.ValidatorFunc[T]) error {
	var wg sync.WaitGroup
	successChan := make(chan struct{})
	errChan := make(chan error, len(values))
	validationErr := hvalid.NewValidationError(v.FieldName)

	for i, value := range values {
		wg.Add(1)
		go func(index int, val T) {
			defer wg.Done()
			if err := validator(val); err == nil {
				select {
				case successChan <- struct{}{}:
				default:
				}
			} else {
				errChan <- fmt.Errorf("element[%d]: %v", index, err)
			}
		}(i, value)
	}

	// 等待所有验证完成或第一个成功的结果
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
}
