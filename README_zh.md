<div align="center">

# hvalid

| [English](README.md) | 中文 |
| --- | --- |

`hvalid` 是一个用 Go 编写的轻量级验证库。它定义了规则类型，并提供一组通用规则，帮助开发者快速完成数据校验。

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)
[![codecov](https://codecov.io/gh/lyonnee/hvalid/graph/badge.svg?token=2LT7V180W9)](https://codecov.io/gh/lyonnee/hvalid)

</div>

## 特性

- **支持泛型。** 可以校验数值、文本、切片、map、时间，以及 `type Age int` 这样的定义类型。结构体由你为它写的规则来校验。
- **易于使用。** 字段名只传给 `Validate` 一次，然后再传入规则。规则本身不携带字段名。
- **可扩展。** 自定义规则和内置规则是同一种类型，可以一起传给 `Validate`。
- **错误信息友好。** 校验失败时返回一条错误。其中有字段名，以及全部失败原因，方便定位问题。

校验失败时的文本像这样：

```text
name: value length too short
count: must be greater than or equal to 1; must be less than or equal to 10
```

## 安装

```bash
go get github.com/lyonnee/hvalid
```

## 校验一个值

```go
import "github.com/lyonnee/hvalid"

err := hvalid.Validate("name", "hi", hvalid.MinLen[string](3))
// name: value length too short

err = hvalid.Validate("count", 10, hvalid.Min(5), hvalid.Max(15))
```

`Validate` 会跑完每一条规则，原因之间用 `"; "` 连接。直接调用 `hvalid.Min(5)(4)` 只会得到 `must be greater than or equal to 5`，因为这里没有字段名。

定义类型走同一批规则：

```go
type Age int

err := hvalid.Validate("age", Age(17), hvalid.Min[Age](18))
```

## 校验结构体

为结构体写一条规则，每个字段调用一次 `Validate`。把这条错误原样返回，字段名才不会被再包一层。

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

err := UserRule()(User{Name: "张三", Age: 20})
```

## 自定义规则

手写规则和内置规则是同一种类型。它返回的错误文本会成为其中一条原因。

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

## 规则

| 规则 | 适用的值 |
| --- | --- |
| `Min`、`Max`、`Range`、`Positive`、`Negative` | 数值，含数值定义类型 |
| `MinLen`、`MaxLen`、`MinRunes`、`MaxRunes` | `string`、`[]byte`，以及文本定义类型。`MinRunes` 与 `MaxRunes` 只用于 `string` |
| `MinItems`、`MaxItems`、`NotEmpty`、`Empty`、`Contains` | 切片 |
| `MinSize`、`MaxSize`、`MapNotEmpty`、`MapEmpty`、`HasKey`、`NoKey` | map |
| `IsTrue`、`IsFalse` | `bool` |
| `Before`、`After`、`Equal`、`Between` | `time.Time` |
| `ContainsStr`、`Regexp`、`ASCII`、`PrintableASCII`、`Alpha`、`Alphanumeric`、`Hex`、`Digits`、`Integer`、`Decimal`、`UTFLetter`、`UTFLetterNumeric`、`UTFNumeric`、`UTFDigit`、`Lower`、`Upper`、`HasLower`、`HasUpper`、`HasSpace`、`SpaceOnly`、`Blank`、`NotBlank`、`OneOf` | `string`，以及文本定义类型 |
| `ContainsBytes` | `[]byte` |
| `Eq` | 任意值 |

`MinLen` 用于文本。切片长度用 `MinItems`。

UUID、ULID、JSON、Base64、data URI、MAC、DNS 名、主机、端口、MongoDB id、JWT、语义化版本、颜色、纬度、经度、RFC3339、Unix 时间、`YYYYMMDD`、ISBN、ISO 3166-1 alpha-2 和 ISO 4217 在 `validators/format`。这些规则不带字段名：

```go
import "github.com/lyonnee/hvalid/validators/format"

err := hvalid.Validate("country", "CN", format.ISO3166Alpha2[string]())
```

邮箱、IPv4、IPv6 和 URL 在 `validators/primitive`，自身已经带字段名，直接调用：

```go
import "github.com/lyonnee/hvalid/validators/primitive"

err := primitive.NewStringValidator("email").IsEmail()("user@example.com")
```

## 测试

```bash
go test ./...
```

## 许可证

MIT。见 [LICENSE](LICENSE)。
