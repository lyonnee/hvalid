package hval

type Validator[T any] interface {
	Validate(field T) error
}

type ValidatorFunc[T any] func(field T) error

func (fn ValidatorFunc[T]) Validate(field T) error {
	return fn(field)
}

func Validate[T any](field T, validators ...Validator[T]) error {
	for _, v := range validators {
		if err := v.Validate(field); err != nil {
			return err
		}
	}

	return nil
}
