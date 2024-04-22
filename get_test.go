package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunc(t *testing.T) {
	var numInf interface{}
	numInf = int(17)

	_, err := Get[int](numInf, Min(18), Max(30))
	assert.ErrorContains(t, err, "value is too small")
}
