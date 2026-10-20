# HValid Validators

规则分三处：

- package `hvalid`：长度、符文长度、数值、切片、map、布尔、时间、相等，以及字符串形态（`ASCII`、`Alpha`、`OneOf` 等）。字段名在 `Validate` 上给一次。
- `format/`：UUID、ULID、JSON、Base64、data URI、MAC、DNS 名、主机、端口、MongoDB id、JWT、语义化版本、颜色、经纬度、时间格式、ISBN、ISO 3166-1 alpha-2、ISO 4217。规则不携带字段名。
- `primitive/`：邮箱、IPv4、IPv6、URL。构造器记住字段名，这是既有例外。
- `common/`：信用卡、邮箱、身份证、IP、密码、电话、邮编、URL 的业务规则。构造器同样记住字段名，这是既有例外。`common` 不依赖 `primitive`。

## 目录结构

- `primitive/`
  - `string.go`: 邮箱、IPv4、IPv6、URL
- `format/`
  - 字符串格式规则，返回 `hvalid.ValidatorFunc`
- `common/`
  - `creditcard.go`: 信用卡号
  - `email.go`: 电子邮件
  - `idcard.go`: 身份证号
  - `ip.go`: IP 地址
  - `password.go`: 密码
  - `phone.go`: 电话号码
  - `postcode.go`: 邮政编码
  - `url.go`: URL

## 使用示例

```go
import (
    "github.com/lyonnee/hvalid"
    "github.com/lyonnee/hvalid/validators/common"
    "github.com/lyonnee/hvalid/validators/format"
)

err := hvalid.Validate("name", "John", hvalid.MinLen[string](1), hvalid.Alpha[string]())
err = hvalid.Validate("country", "CN", format.ISO3166Alpha2[string]())

emailValidator := common.NewEmailValidator("email")
err = emailValidator.Validate()("user@example.com")
```
