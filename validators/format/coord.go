package format

import (
	"errors"
	"math"
	"strconv"
)

func decimalRange(text string, min, max float64, msg string) error {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) || value < min || value > max {
		return errors.New(msg)
	}
	return nil
}
