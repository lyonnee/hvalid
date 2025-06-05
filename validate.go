// Package core provides the core validation interface and implementation
package hvalid

// ValidatorFunc is the function signature for validation functions
type Validator[T any] interface {
	Validate(field T) error
}

type ValidatorFunc[T any] func(field T) error

func (fn ValidatorFunc[T]) Validate(field T) error {
	return fn(field)
}

// Validate 验证字段
func Validate[T any](field T, validators ...ValidatorFunc[T]) error {
	var validationErr *ValidationError

	for _, v := range validators {
		if err := v(field); err != nil {
			if validationErr == nil {
				validationErr = NewValidationError("field")
			}
			validationErr.AddError(err.Error())
		}
	}

	if validationErr != nil && validationErr.HasError() {
		return validationErr
	}

	return nil
}
