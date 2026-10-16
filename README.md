<div align="center">
</br>

# hvalid

| English | [中文](README_zh.md) |
| --- | --- |

`hvalid` is a lightweight validation library written in Go language. It provides a custom validator interface and a series of common validation functions to help developers quickly implement data validation.

</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)
[![codecov](https://codecov.io/gh/lyonnee/hvalid/graph/badge.svg?token=2LT7V180W9)](https://codecov.io/gh/lyonnee/hvalid)

## Features
- Generic support: Can validate any type of data, including basic types, structs, slices, etc.
- Easy to use: Offers a concise API for developers to quickly perform parameter validation.
- Extensible: Allows custom validation rules to meet different validation needs.
- Friendly error messages: Returns clear error messages when validation fails, making it easy for developers to locate issues.

## Installation
Install using the `go get` command:
```bash
go get github.com/lyonnee/hvalid
```

### Usage Examples

#### Basic Type Validation

```go
import (
	"github.com/lyonnee/hvalid"
)

func main() {
	// Validate string length
	err := hvalid.Validate("name", "hello", hvalid.MinLen[string](3))
	if err != nil {
		// Handle error
	}

	// Validate number range
	err = hvalid.Validate("count", 10, hvalid.Min(5), hvalid.Max(15))
	if err != nil {
		// Handle error
	}
}
```

#### Struct Validation

```go
type User struct {
	Name string
	Age  int
}

func UserValidator() hvalid.ValidatorFunc[User] {
	return func(user User) error {
		if err := hvalid.Validate("name", user.Name, hvalid.MinLen[string](1)); err != nil {
			return err
		}
		return hvalid.Validate("age", user.Age, hvalid.Min(18))
	}
}

func main() {
	user := User{
		Name: "Zhang San",
		Age:  20,
	}

	err := UserValidator()(user)
	if err != nil {
		// Handle error
	}
}
```

#### Custom Validation Rules

```go
import "errors"

func IsPositive() hvalid.ValidatorFunc[int] {
	return func(num int) error {
		if num <= 0 {
			return errors.New("The number must be positive")
		}
		return nil
	}
}

func main() {
	err := hvalid.Validate("n", 10, IsPositive())
	if err != nil {
		// Handle error
	}
}
```

## Testing
The project includes unit tests, run all tests with the `go test` command:
```bash
go test ./...
```

## Contributing
Issues and pull requests are welcome to improve `hvalid`.

## License
`hvalid` is released under the MIT License. See the [LICENSE](LICENSE) file for more information.
