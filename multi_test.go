package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	var n1 int = 100
	err := Validate[int](n1, Eq[int](100))
	assert.NoError(t, err)

	var n2 int = 1
	err = Validate[int](n2, Eq[int](10, "int value not eq"))
	assert.Error(t, err)
}

func TestRequired(t *testing.T) {
	// int default is 0
	var n1 int
	err := Validate[int](n1, Required[int]())
	assert.NoError(t, err)

	// string default is ""
	var s1 string
	err = Validate[string](s1, Required[string]("s1 is empty"))
	assert.NoError(t, err)

	fn := func() bool {
		return true
	}
	err = Validate[func() bool](fn, Required[func() bool]())
	assert.NoError(t, err)

	var intPtr1 *int
	err = Validate[*int](intPtr1, Required[*int]("has empty pointer"))
	assert.ErrorContains(t, err, "has empty pointer")

	var a int = 1
	var intPtr2 *int = &a
	err = Validate[*int](intPtr2, Required[*int]())
	assert.NoError(t, err)

	var inf interface{}
	err = Validate[interface{}](inf, Required[interface{}]())
	assert.Error(t, err)
}
