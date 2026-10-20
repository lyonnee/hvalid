package hvalid

import (
	"errors"
	"regexp"
	"unicode"
)

var decimalPattern = regexp.MustCompile(`^[+-]?(?:[0-9]+\.[0-9]+|[0-9]+|\.[0-9]+)(?:[eE][+-]?[0-9]+)?$`)

// ASCII 要求非空，且每个字节都小于 128。
func ASCII[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must be ASCII", func(b byte) bool { return b < 128 })
}

// PrintableASCII 要求非空，且每个字节都在 0x20 到 0x7E。
func PrintableASCII[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must be printable ASCII", func(b byte) bool { return b >= 0x20 && b <= 0x7E })
}

// Alpha 要求非空，且只含 ASCII 字母。
func Alpha[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must contain only ASCII letters", isASCIILetter)
}

// Alphanumeric 要求非空，且只含 ASCII 字母和数字。
func Alphanumeric[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must contain only ASCII letters and digits", func(b byte) bool {
		return isASCIILetter(b) || isASCIIDigit(b)
	})
}

// Hex 要求非空，且只含十六进制字符。
func Hex[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must be hexadecimal", isASCIIHex)
}

// Digits 要求非空，且只含 ASCII 数字。
func Digits[T ~string]() ValidatorFunc[T] {
	return bytesRule[T]("must contain only ASCII digits", isASCIIDigit)
}

// Integer 要求整段是可选符号加十进制整数。
func Integer[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		if !isInteger(string(value)) {
			return errors.New("must be an integer")
		}
		return nil
	}
}

// Decimal 要求整段是十进制数，允许符号、小数和指数。
func Decimal[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		if !decimalPattern.MatchString(string(value)) {
			return errors.New("must be a decimal number")
		}
		return nil
	}
}

// UTFLetter 要求非空，且每个符文都是字母。
func UTFLetter[T ~string]() ValidatorFunc[T] {
	return runesRule[T]("must contain only letters", unicode.IsLetter)
}

// UTFLetterNumeric 要求非空，且每个符文是字母或数字。
func UTFLetterNumeric[T ~string]() ValidatorFunc[T] {
	return runesRule[T]("must contain only letters and numbers", func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
}

// UTFNumeric 要求非空，且每个符文都是数字。
func UTFNumeric[T ~string]() ValidatorFunc[T] {
	return runesRule[T]("must contain only numbers", unicode.IsNumber)
}

// UTFDigit 要求非空，且每个符文都是十进制数字。
func UTFDigit[T ~string]() ValidatorFunc[T] {
	return runesRule[T]("must contain only Unicode digits", unicode.IsDigit)
}

// Lower 要求至少一个字母，且每个字母都是小写。
func Lower[T ~string]() ValidatorFunc[T] {
	return letterCase[T](true, "must be lower case")
}

// Upper 要求至少一个字母，且每个字母都是大写。
func Upper[T ~string]() ValidatorFunc[T] {
	return letterCase[T](false, "must be upper case")
}

// HasLower 要求至少一个小写字母。
func HasLower[T ~string]() ValidatorFunc[T] {
	return hasLetterCase[T](true, "must contain a lower case letter")
}

// HasUpper 要求至少一个大写字母。
func HasUpper[T ~string]() ValidatorFunc[T] {
	return hasLetterCase[T](false, "must contain an upper case letter")
}

// HasSpace 要求至少一个空白符文。
func HasSpace[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		for _, r := range string(value) {
			if unicode.IsSpace(r) {
				return nil
			}
		}
		return errors.New("must contain whitespace")
	}
}

// SpaceOnly 要求非空，且每个符文都是空白。
func SpaceOnly[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		if value == "" {
			return errors.New("must contain only whitespace")
		}
		for _, r := range string(value) {
			if !unicode.IsSpace(r) {
				return errors.New("must contain only whitespace")
			}
		}
		return nil
	}
}

// Blank 要求空串。
func Blank[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		if value != "" {
			return errors.New("must be blank")
		}
		return nil
	}
}

// NotBlank 要求非空串。
func NotBlank[T ~string]() ValidatorFunc[T] {
	return func(value T) error {
		if value == "" {
			return errors.New("must not be blank")
		}
		return nil
	}
}

func bytesRule[T ~string](msg string, ok func(byte) bool) ValidatorFunc[T] {
	return func(value T) error {
		if value == "" {
			return errors.New(msg)
		}
		for i := 0; i < len(value); i++ {
			if !ok(value[i]) {
				return errors.New(msg)
			}
		}
		return nil
	}
}

func runesRule[T ~string](msg string, ok func(rune) bool) ValidatorFunc[T] {
	return func(value T) error {
		if value == "" {
			return errors.New(msg)
		}
		for _, r := range string(value) {
			if !ok(r) {
				return errors.New(msg)
			}
		}
		return nil
	}
}

func letterCase[T ~string](wantLower bool, msg string) ValidatorFunc[T] {
	return func(value T) error {
		letters := 0
		for _, r := range string(value) {
			if !unicode.IsLetter(r) {
				continue
			}
			letters++
			if wantLower && !unicode.IsLower(r) {
				return errors.New(msg)
			}
			if !wantLower && !unicode.IsUpper(r) {
				return errors.New(msg)
			}
		}
		if letters == 0 {
			return errors.New(msg)
		}
		return nil
	}
}

func hasLetterCase[T ~string](wantLower bool, msg string) ValidatorFunc[T] {
	return func(value T) error {
		for _, r := range string(value) {
			if wantLower && unicode.IsLower(r) {
				return nil
			}
			if !wantLower && unicode.IsUpper(r) {
				return nil
			}
		}
		return errors.New(msg)
	}
}

func isInteger(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	if s[0] == '+' || s[0] == '-' {
		i++
	}
	if i == len(s) {
		return false
	}
	for ; i < len(s); i++ {
		if !isASCIIDigit(s[i]) {
			return false
		}
	}
	return true
}

func isASCIILetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func isASCIIDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isASCIIHex(b byte) bool {
	return isASCIIDigit(b) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}
