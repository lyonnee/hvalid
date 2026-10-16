package hvalid

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func requireErr(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
	if err.Error() != want {
		t.Fatalf("got %q, want %q", err.Error(), want)
	}
}

func requireNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMinLen(t *testing.T) {
	requireNil(t, MinLen[string](3)("abc"))
	requireErr(t, MinLen[string](3)("ab"), "value length too short")
	requireNil(t, MaxLen[string](3)("abc"))
	requireErr(t, MaxLen[string](3)("abcd"), "value length too long")

	requireNil(t, MinLen[[]byte](1)([]byte{1}))
	requireErr(t, MinLen[[]byte](1)(nil), "value length too short")

	type Name string
	requireErr(t, MinLen[Name](1)(Name("")), "value length too short")

	requireErr(t, Validate("name", "ab", MinLen[string](3)), "name: value length too short")
}

func TestNumber(t *testing.T) {
	requireNil(t, Min(5)(5))
	requireErr(t, Min(5)(4), "must be greater than or equal to 5")
	requireNil(t, Max(5)(5))
	requireErr(t, Max(5)(6), "must be less than or equal to 5")
	requireNil(t, Range(1, 3)(1))
	requireNil(t, Range(1, 3)(3))
	requireErr(t, Range(1, 3)(0), "must be between 1 and 3")
	requireNil(t, Positive[int]()(1))
	requireErr(t, Positive[int]()(0), "must be positive")
	requireNil(t, Negative[int]()(-1))
	requireErr(t, Negative[int]()(0), "must be negative")

	type Age int
	requireNil(t, Min[Age](18)(18))
	requireErr(t, Validate("age", Age(17), Min[Age](18)), "age: must be greater than or equal to 18")
}

func TestSlice(t *testing.T) {
	requireNil(t, MinItems[int](2)([]int{1, 2}))
	requireErr(t, MinItems[int](1)([]int(nil)), "length must be at least 1")
	requireNil(t, MaxItems[int](1)([]int{1}))
	requireErr(t, MaxItems[int](1)([]int{1, 2}), "length must be at most 1")
	requireNil(t, NotEmpty[int]()([]int{1}))
	requireErr(t, NotEmpty[int]()([]int{}), "must not be empty")
	requireNil(t, Empty[int]()([]int{}))
	requireErr(t, Empty[int]()([]int{1}), "must be empty")
	requireNil(t, Contains(2)([]int{1, 2}))
	requireErr(t, Contains(3)([]int{1, 2}), "must contain the element")
}

func TestMap(t *testing.T) {
	got := map[string]int{"a": 1}
	requireNil(t, MinSize[string, int](1)(got))
	requireErr(t, MinSize[string, int](1)(map[string]int{}), "must have at least 1 entries")
	requireNil(t, MaxSize[string, int](1)(got))
	requireErr(t, MaxSize[string, int](0)(got), "must have at most 0 entries")
	requireNil(t, MapNotEmpty[string, int]()(got))
	requireErr(t, MapNotEmpty[string, int]()(map[string]int{}), "must not be empty")
	requireNil(t, MapEmpty[string, int]()(map[string]int{}))
	requireErr(t, MapEmpty[string, int]()(got), "must be empty")
	requireNil(t, HasKey[string, int]("a")(got))
	requireErr(t, HasKey[string, int]("b")(got), "must contain key: b")
	requireNil(t, NoKey[string, int]("b")(got))
	requireErr(t, NoKey[string, int]("a")(got), "must not contain key: a")
}

func TestBool(t *testing.T) {
	requireNil(t, IsTrue()(true))
	requireErr(t, IsTrue()(false), "must be true")
	requireNil(t, IsFalse()(false))
	requireErr(t, IsFalse()(true), "must be false")
}

func TestTime(t *testing.T) {
	start := time.Date(2026, 10, 16, 12, 0, 0, 0, time.UTC)
	mid := start.Add(time.Hour)
	end := start.Add(2 * time.Hour)

	requireNil(t, Before(mid)(start))
	requireErr(t, Before(start)(start), fmt.Sprintf("must be before %v", start))
	requireNil(t, After(start)(mid))
	requireErr(t, After(mid)(mid), fmt.Sprintf("must be after %v", mid))
	requireNil(t, Equal(start)(start))
	requireErr(t, Equal(start)(mid), fmt.Sprintf("must be equal to %v", start))
	requireNil(t, Between(start, end)(start))
	requireNil(t, Between(start, end)(end))
	requireErr(t, Between(start, end)(end.Add(time.Second)), fmt.Sprintf("must be between %v and %v", start, end))
}

func TestBytesAndPattern(t *testing.T) {
	requireNil(t, ContainsBytes([]byte("b"))([]byte("ab")))
	requireErr(t, ContainsBytes([]byte("c"))([]byte("ab")), "must contain the sub byte slice")
	requireNil(t, ContainsStr("b")("ab"))
	requireErr(t, ContainsStr("c")("ab"), "must contain the sub string")
	requireNil(t, Regexp("^a+$")("aa"))
	requireErr(t, Regexp("^a+$")("b"), "must match the required pattern")
}

func TestEq(t *testing.T) {
	requireNil(t, Eq(1)(1))
	requireErr(t, Eq(1)(2), "the two values are not equal")
}

func TestValidateCollectsReasons(t *testing.T) {
	err := Validate("n", 0, Min(1), Max(-1), ValidatorFunc[int](func(int) error {
		return errors.New("too young")
	}))
	requireErr(t, err, "n: must be greater than or equal to 1; must be less than or equal to -1; too young")

	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("Validate should return a validation error")
	}
	if ve.Field != "n" {
		t.Fatalf("field = %q", ve.Field)
	}

	if err := Min(1)(0); err == nil {
		t.Fatal("expected an error")
	} else if errors.As(err, &ve) {
		t.Fatal("a rule alone should not carry a field")
	}

	requireNil(t, Validate("n", 1))
}
