package hval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContains(t *testing.T) {
	var s1 string = "hello,world"
	substr := "llo"
	err := Validat[string](s1, ContainsStr(substr))
	assert.NoError(t, err)

	var s2 string = "lyon.nee@outlook.com"
	substr = "nee"
	err = Validat[string](s2, ContainsStr(substr), Email())
	assert.NoError(t, err)
}
