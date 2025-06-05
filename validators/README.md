# HValid Validators

HValid 验证器库提供全面的数据验证功能。

## 目录结构

- `primitive/`: 基础验证器，提供基本的验证功能
  - `boolean.go`: 布尔值验证器
  - `bytes.go`: 字节切片验证器
  - `number.go`: 数字验证器
  - `slice.go`: 切片验证器
  - `text.go`: 文本验证器
  - `time.go`: 时间验证器

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
    "github.com/lyonnee/hvalid/validators/primitive"
    "github.com/lyonnee/hvalid/validators/common"
)

// 使用基础验证器
textValidator := primitive.NewTextValidator("name")
err := textValidator.Required()("John")

// 使用通用验证器
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