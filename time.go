package hvalid

import (
	"fmt"
	"time"
)

// Before 要求时刻早于 t。
func Before(t time.Time) ValidatorFunc[time.Time] {
	return func(value time.Time) error {
		if !value.Before(t) {
			return fmt.Errorf("must be before %v", t)
		}
		return nil
	}
}

// After 要求时刻晚于 t。
func After(t time.Time) ValidatorFunc[time.Time] {
	return func(value time.Time) error {
		if !value.After(t) {
			return fmt.Errorf("must be after %v", t)
		}
		return nil
	}
}

// Equal 要求时刻与 t 相同。
func Equal(t time.Time) ValidatorFunc[time.Time] {
	return func(value time.Time) error {
		if !value.Equal(t) {
			return fmt.Errorf("must be equal to %v", t)
		}
		return nil
	}
}

// Between 要求时刻落在 start 与 end 之间，含两端。
func Between(start, end time.Time) ValidatorFunc[time.Time] {
	return func(value time.Time) error {
		if value.Before(start) || value.After(end) {
			return fmt.Errorf("must be between %v and %v", start, end)
		}
		return nil
	}
}
