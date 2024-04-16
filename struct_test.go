package hvalid

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Student struct {
	Name  string
	Email string
	Age   int
	Score int
}

func StudentValidator() ValidatorFunc[Student] {
	return ValidatorFunc[Student](func(data Student) error {
		if data.Age < 0 {
			return errors.New("学生的年龄不能小于 0")
		}

		if data.Score > 100 {
			return errors.New("课程成绩不能大于 100")
		}

		return Validate[string](data.Email, IsEmail("Not email address"))
	})
}

func TestStruct(t *testing.T) {
	var s1 = Student{
		Name:  "张三",
		Email: "zhangsan@gmail.com",
		Age:   18,
		Score: 110,
	}
	err1 := Validate[Student](s1, StudentValidator())
	assert.ErrorContains(t, err1, "课程成绩不能大于 100")

	var s2 = Student{
		Name:  "李四",
		Email: "lisi.gmail.com",
		Age:   19,
		Score: 59,
	}
	err2 := Validate[Student](s2, StudentValidator())
	assert.ErrorContains(t, err2, "Not email address")

	var s3 = Student{
		Name:  "lyonnee",
		Email: "lyon.nee@gmail.com",
		Age:   20,
		Score: 99,
	}
	err3 := Validate[Student](s3, StudentValidator())
	assert.NoError(t, err3)
}
