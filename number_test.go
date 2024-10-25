package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type customInt int

func TestNumberMin(t *testing.T) {
	var n1 int = 256
	err := Validate[int](n1, Min(1000, ""))
	assert.Error(t, err)

	var n2 uint = 256
	err = Validate[uint](n2, Min(uint(255), ""))
	assert.NoError(t, err)

	assert.Error(t, Validate[customInt](customInt(256), Min(customInt(1000), "")))
}

func TestNumberMax(t *testing.T) {
	var n1 int = 256
	err := Validate[int](n1, Max(255))
	assert.Error(t, err)

	var n2 uint = 256
	err = Validate[uint](n2, Max(uint(1000)))
	assert.NoError(t, err)

	assert.Error(t, Validate[customInt](customInt(1000), Max(customInt(256), "")))
}

func TestNumberLimit(t *testing.T) {
	var n1 int = 100
	err := Validate[int](n1, Min(110, "The num is less than the minimum value"), Max(200))
	assert.ErrorContains(t, err, "The num is less than the minimum value")
	var n2 uint = 100
	err = Validate[uint](n2, Min(uint(1)), Max(uint(99), "The num is greater than the maximum value, maximum value"))
	assert.ErrorContains(t, err, "The num is greater than the maximum value, maximum value")

	var n3 float32 = 0.12
	err = Validate[float32](n3, Min(float32(0.1)), Max(float32(1.0)))
	assert.NoError(t, err)
}
