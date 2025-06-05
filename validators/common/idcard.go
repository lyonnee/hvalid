package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// IDCardValidator 身份证号验证器
type IDCardValidator struct {
	FieldName string // 字段名称
}

// NewIDCardValidator 创建一个新的身份证号验证器
func NewIDCardValidator(fieldName string) *IDCardValidator {
	return &IDCardValidator{
		FieldName: fieldName,
	}
}

// Validate 验证身份证号格式
func (v *IDCardValidator) Validate() func(string) error {
	return func(s string) error {
		// 验证长度
		if len(s) != 18 {
			return fmt.Errorf("身份证号长度必须为18位")
		}

		// 验证前17位是否都是数字
		for i := 0; i < 17; i++ {
			if s[i] < '0' || s[i] > '9' {
				return fmt.Errorf("身份证号前17位必须都是数字")
			}
		}

		// 验证最后一位是否为数字或X
		lastChar := s[17]
		if (lastChar < '0' || lastChar > '9') && lastChar != 'X' && lastChar != 'x' {
			return fmt.Errorf("身份证号最后一位必须是数字或X")
		}

		return nil
	}
}

// ValidateAreaCode 验证身份证号地区码
func (v *IDCardValidator) ValidateAreaCode(validAreaCodes map[string]string) func(string) error {
	return func(s string) error {
		// 验证地区码（前6位）
		areaCode := s[:6]
		if _, ok := validAreaCodes[areaCode]; !ok {
			return fmt.Errorf("无效的地区码")
		}
		return nil
	}
}

// ValidateBirthDate 验证身份证号出生日期
func (v *IDCardValidator) ValidateBirthDate(minAge, maxAge int) func(string) error {
	return func(s string) error {
		// 提取出生日期（第7-14位）
		birthDate := s[6:14]
		year, _ := strconv.Atoi(birthDate[:4])
		month, _ := strconv.Atoi(birthDate[4:6])
		day, _ := strconv.Atoi(birthDate[6:8])

		// 验证日期是否有效
		birthTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		if birthTime.Year() != year || birthTime.Month() != time.Month(month) || birthTime.Day() != day {
			return fmt.Errorf("无效的出生日期")
		}

		// 验证年龄范围
		age := time.Now().Year() - year
		if age < minAge || age > maxAge {
			return fmt.Errorf("年龄超出有效范围")
		}

		return nil
	}
}

// ValidateCheckCode 验证身份证号校验码
func (v *IDCardValidator) ValidateCheckCode() func(string) error {
	return func(s string) error {
		// 加权因子
		weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		// 校验码对应值
		checkCodes := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

		// 计算校验码
		sum := 0
		for i := 0; i < 17; i++ {
			num, _ := strconv.Atoi(string(s[i]))
			sum += num * weights[i]
		}
		checkCode := checkCodes[sum%11]

		// 验证校验码
		if checkCode != strings.ToUpper(s[17:])[0] {
			return fmt.Errorf("校验码错误")
		}

		return nil
	}
}
