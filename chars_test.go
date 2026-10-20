package hvalid

import (
	"errors"
	"testing"
)

func TestChars(t *testing.T) {
	type Name string

	requireNil(t, ASCII[string]()("A~"))
	requireErr(t, ASCII[string]()(""), "must be ASCII")
	requireErr(t, ASCII[string]()("你"), "must be ASCII")
	requireNil(t, ASCII[Name]()(Name("ok")))

	requireNil(t, PrintableASCII[string]()(" ~"))
	requireErr(t, PrintableASCII[string]()(""), "must be printable ASCII")
	requireErr(t, PrintableASCII[string]()("\n"), "must be printable ASCII")
	requireErr(t, PrintableASCII[string]()("\x7f"), "must be printable ASCII")

	requireNil(t, Alpha[string]()("Az"))
	requireErr(t, Alpha[string]()(""), "must contain only ASCII letters")
	requireErr(t, Alpha[string]()("A1"), "must contain only ASCII letters")

	requireNil(t, Alphanumeric[string]()("A1z"))
	requireErr(t, Alphanumeric[string]()(""), "must contain only ASCII letters and digits")
	requireErr(t, Alphanumeric[string]()("A_1"), "must contain only ASCII letters and digits")

	requireNil(t, Hex[string]()("0aF"))
	requireErr(t, Hex[string]()(""), "must be hexadecimal")
	requireErr(t, Hex[string]()("0g"), "must be hexadecimal")

	requireNil(t, Digits[string]()("09"))
	requireErr(t, Digits[string]()(""), "must contain only ASCII digits")
	requireErr(t, Digits[string]()("0a"), "must contain only ASCII digits")

	requireNil(t, Integer[string]()("0"))
	requireNil(t, Integer[string]()("-1"))
	requireNil(t, Integer[string]()("+1"))
	requireErr(t, Integer[string]()(""), "must be an integer")
	requireErr(t, Integer[string]()("-"), "must be an integer")
	requireErr(t, Integer[string]()("+"), "must be an integer")
	requireErr(t, Integer[string]()("1.0"), "must be an integer")
	requireErr(t, Integer[string]()("1a"), "must be an integer")

	requireNil(t, Decimal[string]()(".5"))
	requireNil(t, Decimal[string]()("1e-2"))
	requireNil(t, Decimal[string]()("+1.5"))
	requireNil(t, Decimal[string]()("1E+2"))
	requireNil(t, Decimal[string]()("10"))
	requireErr(t, Decimal[string]()(""), "must be a decimal number")
	requireErr(t, Decimal[string]()("1."), "must be a decimal number")
	requireErr(t, Decimal[string]()("NaN"), "must be a decimal number")
	requireErr(t, Decimal[string]()("Inf"), "must be a decimal number")
	requireErr(t, Decimal[string]()("0x1"), "must be a decimal number")
	requireErr(t, Decimal[string]()("1e"), "must be a decimal number")

	requireNil(t, UTFLetter[string]()("你好"))
	requireErr(t, UTFLetter[string]()(""), "must contain only letters")
	requireErr(t, UTFLetter[string]()("a1"), "must contain only letters")

	requireNil(t, UTFLetterNumeric[string]()("a1"))
	requireNil(t, UTFLetterNumeric[string]()("½"))
	requireErr(t, UTFLetterNumeric[string]()(""), "must contain only letters and numbers")
	requireErr(t, UTFLetterNumeric[string]()("a-"), "must contain only letters and numbers")

	requireNil(t, UTFNumeric[string]()("123"))
	requireNil(t, UTFNumeric[string]()("½"))
	requireErr(t, UTFNumeric[string]()(""), "must contain only numbers")
	requireErr(t, UTFNumeric[string]()("a"), "must contain only numbers")

	requireNil(t, UTFDigit[string]()("123"))
	requireNil(t, UTFDigit[string]()("٣"))
	requireErr(t, UTFDigit[string]()(""), "must contain only Unicode digits")
	requireErr(t, UTFDigit[string]()("½"), "must contain only Unicode digits")

	requireNil(t, Lower[string]()("a1"))
	requireNil(t, Lower[string]()("ä"))
	requireErr(t, Lower[string]()(""), "must be lower case")
	requireErr(t, Lower[string]()("1"), "must be lower case")
	requireErr(t, Lower[string]()("aB"), "must be lower case")
	requireErr(t, Lower[string]()("Ä"), "must be lower case")

	requireNil(t, Upper[string]()("A1"))
	requireNil(t, Upper[string]()("Ä"))
	requireErr(t, Upper[string]()(""), "must be upper case")
	requireErr(t, Upper[string]()("1"), "must be upper case")
	requireErr(t, Upper[string]()("Ab"), "must be upper case")

	requireNil(t, HasLower[string]()("A1a"))
	requireErr(t, HasLower[string]()("A1"), "must contain a lower case letter")
	requireNil(t, HasUpper[string]()("aA"))
	requireErr(t, HasUpper[string]()("a1"), "must contain an upper case letter")

	requireNil(t, HasSpace[string]()("a b"))
	requireNil(t, HasSpace[string]()("\t"))
	requireErr(t, HasSpace[string]()("ab"), "must contain whitespace")

	requireNil(t, SpaceOnly[string]()(" \t\n"))
	requireErr(t, SpaceOnly[string]()(""), "must contain only whitespace")
	requireErr(t, SpaceOnly[string]()(" a"), "must contain only whitespace")

	requireNil(t, Blank[string]()(""))
	requireErr(t, Blank[string]()("a"), "must be blank")
	requireNil(t, NotBlank[string]()("a"))
	requireErr(t, NotBlank[string]()(""), "must not be blank")

	err := ASCII[string]()("")
	var ve *ValidationError
	if !errors.As(err, &ve) {
		return
	}
	t.Fatal("a rule alone should not carry a field")
}
