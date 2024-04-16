package hval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrBytesLen(t *testing.T) {
	var str1 = "i am developer"
	err := Validat[string](str1, MinLen[string](1), MaxLen[string](10))
	assert.Error(t, err)

	var str2 = "golang"
	err = Validat[string](str2, MinLen[string](2), MaxLen[string](6))
	assert.NoError(t, err)

	var bytes = []byte("bytes value")
	err = Validat[[]byte](bytes, MinLen[[]byte](3), MaxLen[[]byte](10))
	assert.Error(t, err)
}
