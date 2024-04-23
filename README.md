<div align="center">
</br>

# hvalid

| English | [中文](README_zh.md) |
| --- | --- |

`hvalid` is a lightweight validation library written in Go language. It provides a custom validator interface and a series of common validation functions to help developers quickly implement data validation.

</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)

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
	"errors"
	"github.com/lyonnee/hvalid"
)

func main() {
	// Validate string length
	err := hvalid.Validate[string]("hello", hvalid.MinLen[string](3))
	if err != nil {
		// Handle error
	}

	// Validate number range
	err = hvalid.Validate[int](10, hvalid.Min(5), hvalid.Max(15))
	if err != nil {
		// Handle error
	}
}
```

#### Struct Validation

```go
type User struct {
	Name  string
	Email string
	Age   int
}

func UserValidator() hvalid.ValidatorFunc[User] {
	return hvalid.ValidatorFunc[User](func(user User) error {
		if user.Age < 18 {
			return errors.New("Age must be greater than 18")
		}

		return hvalid.Validate[string](user.Email, hvalid.Email())
	})
}

func main() {
	user := User{
		Name:  "Zhang San",
		Email: "zhangsan@example.com",
		Age:   20,
	}

	err := hvalid.Validate[User](user, UserValidator())
	if err != nil {
		// Handle error
	}
}
```

#### Custom Validation Rules

```go
func IsPositive(errMsg ...string) hvalid.ValidatorFunc[int] {
	return hvalid.ValidatorFunc[int](func(num int) error {
		if num <= 0 {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("The number must be positive")
		}
		return nil
	})
}

func main() {
	err := hvalid.Validate[int](10, IsPositive())
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