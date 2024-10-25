package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	customString string
	customBytes  []byte
)

func TestStrBytesLen(t *testing.T) {
	var str1 = "i am developer"
	err := Validate[string](str1, MinLen[string](25, "字符串长度太短"), MaxLen[string](100))
	assert.Error(t, err)

	var str2 = "golang"
	err = Validate[string](str2, MinLen[string](2), MaxLen[string](6))
	assert.NoError(t, err)

	err = Validate[customString](customString("f"), MinLen[customString](25), MaxLen[customString](100))
	assert.Error(t, err)

	var bytes1 = []byte("bytes value")
	err = Validate[[]byte](bytes1, MinLen[[]byte](15, "字节数组长度太短"), MaxLen[[]byte](30))
	assert.Error(t, err)

	var bytes2 = []byte("bytes value....")
	err = Validate[[]byte](bytes2, MinLen[[]byte](1), MaxLen[[]byte](10, "超出数组长度最大值"))
	assert.Error(t, err)

	err = Validate[customBytes](customBytes("custom_bytes"), MinLen[customBytes](25), MaxLen[customBytes](100))
	assert.Error(t, err)
}
