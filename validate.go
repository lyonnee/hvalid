// Package hvalid 判断一个值是否符合调用方给出的规则。
package hvalid

// Validator 是规则的运行方式。
type Validator[T any] interface {
	Validate(field T) error
}

type ValidatorFunc[T any] func(field T) error

func (fn ValidatorFunc[T]) Validate(field T) error {
	return fn(field)
}

// Validate 用 name 运行全部规则，把每条失败原因收成一条错误。
func Validate[T any](name string, value T, rules ...ValidatorFunc[T]) error {
	var validationErr *ValidationError

	for _, rule := range rules {
		if err := rule(value); err != nil {
			if validationErr == nil {
				validationErr = NewValidationError(name)
			}
			validationErr.AddError(err.Error())
		}
	}

	if validationErr != nil && validationErr.HasError() {
		return validationErr
	}

	return nil
}
