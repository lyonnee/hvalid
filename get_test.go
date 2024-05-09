package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNum(t *testing.T) {
	var numInf interface{}
	numInf = int(17)
	_, err := Get[int](numInf, Min(18), Max(30))
	assert.ErrorContains(t, err, "value is too small")

	numInf = float32(0.12)
	_, err = Get[float32](numInf, Min(float32(0.1)), Eq(float32(0.12)))
	assert.NoError(t, err)

	var bytesInf interface{}
	bytesInf = []byte("abcdefg")
	_, err = Get[[]byte](bytesInf, ContainsBytes([]byte("cde")))
	assert.NoError(t, err)

	var strInf interface{}
	strInf = []byte("iamlyon")
	_, err = Get[string](strInf)
	assert.Error(t, err)
}
