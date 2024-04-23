
<div align="center">
</br>

# hvalid

| [English](README.md) | 中文 |
| --- | --- |

`hvalid` 是一个用Go语言编写的轻量级验证库，它自定义验证器的接口，以及提供了一系列通用的验证函数，以帮助开发者快速实现数据验证。
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/hvalid)](https://goreportcard.com/report/github.com/lyonnee/hvalid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/hvalid)

## 特性
- 支持泛型: 可以对任何类型的数据进行校验，包括基本类型、结构体、切片等。
- 易于使用: 提供简洁的 API，方便开发者快速进行参数校验。
- 可扩展: 可以自定义校验规则，满足不同的校验需求。
- 错误信息友好: 校验失败时，会返回清晰的错误信息，方便开发者定位问题。

## 安装
使用`go get`命令安装:
```bash
go get github.com/lyonnee/hvalid
```


### 使用示例

#### 基本类型校验

```go
import (
	"errors"
	"github.com/lyonnee/hvalid"
)

func main() {
	// 校验字符串长度
	err := hvalid.Validate[string]("hello", hvalid.MinLen[string](3))
	if err != nil {
		// 处理错误
	}

	// 校验数字范围
	err = hvalid.Validate[int](10, hvalid.Min(5), hvalid.Max(15))
	if err != nil {
		// 处理错误
	}
}
```

#### 结构体校验

```go
type User struct {
	Name  string
	Email string
	Age   int
}

func UserValidator() hvalid.ValidatorFunc[User] {
	return hvalid.ValidatorFunc[User](func(user User) error {
		if user.Age < 18 {
			return errors.New("年龄必须大于 18 岁")
		}

		return hvalid.Validate[string](user.Email, hvalid.Email())
	})
}

func main() {
	user := User{
		Name:  "张三",
		Email: "zhangsan@example.com",
		Age:   20,
	}

	err := hvalid.Validate[User](user, UserValidator())
	if err != nil {
		// 处理错误
	}
}
```

#### 自定义校验规则

```go
func IsPositive(errMsg ...string) hvalid.ValidatorFunc[int] {
	return hvalid.ValidatorFunc[int](func(num int) error {
		if num <= 0 {
			if len(errMsg) > 0 {
				return errors.New(errMsg[0])
			}
			return errors.New("数字必须为正数")
		}
		return nil
	})
}

func main() {
	err := hvalid.Validate[int](10, IsPositive())
	if err != nil {
		// 处理错误
	}
}
```

## 测试
项目包含单元测试，使用`go test`命令执行所有测试：
```bash
go test ./...
```

## 贡献
欢迎提交问题和拉取请求来改进`hvalid`。

## 许可证
`hvalid`遵循MIT许可证。查看[LICENSE](LICENSE)文件以获取更多信息。