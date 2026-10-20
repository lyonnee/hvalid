package hvalid

import "testing"

func TestOneOf(t *testing.T) {
	type Code string

	requireNil(t, OneOf("b", "a")("a"))
	requireErr(t, OneOf("a", "b")("c"), "must be one of: a, b")
	requireErr(t, OneOf[string]()("a"), "must be one of: ")
	requireNil(t, OneOf[Code]("ab")(Code("ab")))
	requireErr(t, OneOf("A")("a"), "must be one of: A")
}
