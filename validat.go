package hval

type Validator[T any] interface {
	Validat(field T) error
}

type ValidatorFunc[T any] func(field T) error

func (fn ValidatorFunc[T]) Validat(field T) error {
	return fn(field)
}

func Validat[T any](field T, validators ...Validator[T]) error {
	for _, v := range validators {
		if err := v.Validat(field); err != nil {
			return err
		}
	}

	return nil
}
