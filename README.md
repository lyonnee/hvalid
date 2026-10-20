<div align="center">

# hvalid

| English | [中文](README_zh.md) |
| --- | --- |

`hvalid` is a lightweight validation library written in Go. It defines a rule type and ships common rules, so you can check data quickly.

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)
[![codecov](https://codecov.io/gh/lyonnee/hvalid/graph/badge.svg?token=2LT7V180W9)](https://codecov.io/gh/lyonnee/hvalid)

</div>

## Features

- **Generics.** Check numbers, text, slices, maps, times, and defined types such as `type Age int`. A struct is checked by a rule you write for that struct.
- **Small call.** Pass the field name once to `Validate`, then pass the rules. A rule does not carry the field name.
- **Extensible.** A custom rule has the same type as a built-in rule, so both can be passed to `Validate`.
- **Clear errors.** Failure returns one error. It names the field and lists every reason.

A failed check looks like this:

```text
name: value length too short
count: must be greater than or equal to 1; must be less than or equal to 10
```

## Install

```bash
go get github.com/lyonnee/hvalid
```

## Check one value

```go
import "github.com/lyonnee/hvalid"

err := hvalid.Validate("name", "hi", hvalid.MinLen[string](3))
// name: value length too short

err = hvalid.Validate("count", 10, hvalid.Min(5), hvalid.Max(15))
```

`Validate` runs every rule. Reasons are joined with `"; "`. A direct call such as `hvalid.Min(5)(4)` returns only `must be greater than or equal to 5`, because no field name was given.

Defined types use the same rules:

```go
type Age int

err := hvalid.Validate("age", Age(17), hvalid.Min[Age](18))
```

## Check a struct

Write a rule for the struct and call `Validate` once per field. Return that error as-is so the field name stays intact.

```go
type User struct {
	Name string
	Age  int
}

func UserRule() hvalid.ValidatorFunc[User] {
	return func(user User) error {
		if err := hvalid.Validate("name", user.Name, hvalid.MinLen[string](1)); err != nil {
			return err
		}
		return hvalid.Validate("age", user.Age, hvalid.Min(18))
	}
}

err := UserRule()(User{Name: "Zhang San", Age: 20})
```

## Custom rules

A custom rule is the same type as a built-in one. Its error text becomes one reason.

```go
import "errors"

func IsPositive() hvalid.ValidatorFunc[int] {
	return func(n int) error {
		if n <= 0 {
			return errors.New("must be positive")
		}
		return nil
	}
}

err := hvalid.Validate("n", 0, IsPositive())
// n: must be positive
```

## Rules

| Rule | Value |
| --- | --- |
| `Min`, `Max`, `Range`, `Positive`, `Negative` | numbers, including defined numeric types |
| `MinLen`, `MaxLen`, `MinRunes`, `MaxRunes` | `string`, `[]byte`, and defined text types. `MinRunes` and `MaxRunes` are `string` only |
| `MinItems`, `MaxItems`, `NotEmpty`, `Empty`, `Contains` | slices |
| `MinSize`, `MaxSize`, `MapNotEmpty`, `MapEmpty`, `HasKey`, `NoKey` | maps |
| `IsTrue`, `IsFalse` | `bool` |
| `Before`, `After`, `Equal`, `Between` | `time.Time` |
| `ContainsStr`, `Regexp`, `ASCII`, `PrintableASCII`, `Alpha`, `Alphanumeric`, `Hex`, `Digits`, `Integer`, `Decimal`, `UTFLetter`, `UTFLetterNumeric`, `UTFNumeric`, `UTFDigit`, `Lower`, `Upper`, `HasLower`, `HasUpper`, `HasSpace`, `SpaceOnly`, `Blank`, `NotBlank`, `OneOf` | `string`, including defined text types |
| `ContainsBytes` | `[]byte` |
| `Eq` | any value |

`MinLen` is for text. Slice length is `MinItems`.

UUID, ULID, JSON, Base64, data URI, MAC, DNS name, host, port, MongoDB id, JWT, semver, colors, latitude, longitude, RFC3339, Unix time, `YYYYMMDD`, ISBN, ISO 3166-1 alpha-2, and ISO 4217 live in `validators/format`. Those rules do not carry a field name:

```go
import "github.com/lyonnee/hvalid/validators/format"

err := hvalid.Validate("country", "CN", format.ISO3166Alpha2[string]())
```

Email, IPv4, IPv6, and URL checks live in `validators/primitive` and already include a field name, so call them directly:

```go
import "github.com/lyonnee/hvalid/validators/primitive"

err := primitive.NewStringValidator("email").IsEmail()("user@example.com")
```

## Test

```bash
go test ./...
```

## License

MIT. See [LICENSE](LICENSE).
