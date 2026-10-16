# HValid Validators

HValid 验证器库提供全面的数据验证功能。

## 目录结构

- `primitive/`: 字符串格式检查（邮箱、IPv4、IPv6、URL）。长度、数值、切片、map、布尔和时间规则在 package `hvalid`。
  - `string.go`: 邮箱、IPv4、IPv6、URL

- `common/`: 通用验证器，提供常用的验证功能
  - `creditcard.go`: 信用卡号验证器
  - `email.go`: 电子邮件验证器
  - `idcard.go`: 身份证号验证器
  - `ip.go`: IP地址验证器
  - `password.go`: 密码验证器
  - `phone.go`: 电话号码验证器
  - `postcode.go`: 邮政编码验证器
  - `url.go`: URL验证器

## 使用示例

```go
import (
    "github.com/lyonnee/hvalid"
    "github.com/lyonnee/hvalid/validators/common"
)

err := hvalid.Validate("name", "John", hvalid.MinLen[string](1))

emailValidator := common.NewEmailValidator("email")
err = emailValidator.Validate()("user@example.com")
```

## 验证器关系

```
primitive/ (基础验证器)
    ↓
common/ (通用验证器)
```

## 注意事项

1. 所有验证器都支持链式调用
2. 可以根据业务需求扩展验证器
3. 提供全面的数据验证功能 