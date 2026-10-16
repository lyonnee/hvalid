<div align="center">

# hvalid

| [English](README.md) | 中文 |
| --- | --- |

用 Go 写的轻量校验库。

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)
[![codecov](https://codecov.io/gh/lyonnee/hvalid/graph/badge.svg?token=2LT7V180W9)](https://codecov.io/gh/lyonnee/hvalid)

</div>

字段名只传一次。每条规则只判断值是否通过。失败时 `Validate` 返回一条错误，里面有字段名，以及全部原因。

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
| `MinLen`、`MaxLen` | `string`、`[]byte`，以及文本定义类型 |
| `MinItems`、`MaxItems`、`NotEmpty`、`Empty`、`Contains` | 切片 |
| `MinSize`、`MaxSize`、`MapNotEmpty`、`MapEmpty`、`HasKey`、`NoKey` | map |
| `IsTrue`、`IsFalse` | `bool` |
| `Before`、`After`、`Equal`、`Between` | `time.Time` |
| `ContainsStr`、`Regexp` | `string` |
| `ContainsBytes` | `[]byte` |
| `Eq` | 任意值 |

`MinLen` 用于文本。切片长度用 `MinItems`。

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
