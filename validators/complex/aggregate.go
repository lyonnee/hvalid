package complex

import (
	"fmt"

	"github.com/lyonnee/hvalid"
)

// AggregateValidator 聚合验证器结构体
type AggregateValidator[T any] struct {
	FieldName string // 字段名称
}

// NewAggregateValidator 创建聚合验证器
func NewAggregateValidator[T any](fieldName string) *AggregateValidator[T] {
	return &AggregateValidator[T]{
		FieldName: fieldName,
	}
}

// Aggregate 聚合多个验证器的结果
func (v *AggregateValidator[T]) Aggregate(validators ...hvalid.ValidatorFunc[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)
		successCount := 0

		for i, validator := range validators {
			if err := validator(value); err != nil {
				validationErr.AddError(fmt.Sprintf("validator[%d]: %v", i, err))
			} else {
				successCount++
			}
		}

		if successCount == 0 {
			return validationErr
		}
		return nil
	})
}

// AggregateWithWeight 使用权重聚合多个验证器的结果
// 由于 Go 不支持以函数为 map key，这里改为传递一个包含验证器和权重的切片

type WeightedValidator[T any] struct {
	Validator hvalid.ValidatorFunc[T]
	Weight    float64
}

func (v *AggregateValidator[T]) AggregateWithWeight(weightedValidators []WeightedValidator[T]) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)
		totalWeight := 0.0
		successWeight := 0.0

		for i, wv := range weightedValidators {
			if err := wv.Validator(value); err != nil {
				validationErr.AddError(fmt.Sprintf("validator[%d, weight=%.2f]: %v", i, wv.Weight, err))
			} else {
				successWeight += wv.Weight
			}
			totalWeight += wv.Weight
		}

		if totalWeight == 0 || successWeight/totalWeight < 0.5 {
			return validationErr
		}
		return nil
	})
}

// AggregateWithThreshold 使用阈值聚合多个验证器的结果
func (v *AggregateValidator[T]) AggregateWithThreshold(validators []hvalid.ValidatorFunc[T], threshold float64) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		validationErr := hvalid.NewValidationError(v.FieldName)
		successCount := 0

		for i, validator := range validators {
			if err := validator(value); err != nil {
				validationErr.AddError(fmt.Sprintf("validator[%d]: %v", i, err))
			} else {
				successCount++
			}
		}

		successRate := float64(successCount) / float64(len(validators))
		if successRate < threshold {
			return validationErr
		}
		return nil
	})
}

// AggregateWithCustom 使用自定义聚合函数聚合多个验证器的结果
func (v *AggregateValidator[T]) AggregateWithCustom(validators []hvalid.ValidatorFunc[T], aggregate func([]error) error) hvalid.ValidatorFunc[T] {
	return hvalid.ValidatorFunc[T](func(value T) error {
		var errors []error

		for i, validator := range validators {
			if err := validator(value); err != nil {
				errors = append(errors, fmt.Errorf("validator[%d]: %v", i, err))
			}
		}

		return aggregate(errors)
	})
}
