package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContains(t *testing.T) {
	var s1 string = "hello,world"
	substr := "llo"
	err := Validate[string](s1, ContainsStr(substr))
	assert.NoError(t, err)
}

func TestIsEmail(t *testing.T) {
	var e1 string = "lyon.nee@outlook.com"
	err := Validate[string](e1, IsEmail("无效的邮箱地址"))
	assert.NoError(t, err)

	var e2 string = "lyon.neeoutlook.com"
	err = Validate[string](e2, IsEmail("无效的邮箱地址"))
	assert.Error(t, err)
}

func TestIsUrl(t *testing.T) {
	err := Validate[string]("testURL", IsUrl("无效的url"))
	assert.Error(t, err)

	err = Validate[string]("lyon.nee/", IsUrl())
	assert.Error(t, err)

	err = Validate[string]("http://github.com", IsUrl())
	assert.NoError(t, err)

	err = Validate[string]("https://github.com/lyonnee/hvalid", IsUrl())
	assert.NoError(t, err)
}
