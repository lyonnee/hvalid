# Common Validators

本包提供基于基础验证器的通用验证功能。

## 包含的验证器

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
import "github.com/lyonnee/hvalid/validators/common"

// 创建电子邮件验证器
emailValidator := common.NewEmailValidator("email")

// 验证电子邮件格式
err := emailValidator.Validate()("user@example.com")

// 验证电子邮件域名
err = emailValidator.ValidateDomain([]string{"example.com"})("user@example.com")
```

## 注意事项

1. 所有验证器都基于基础验证器构建
2. 每个验证器都提供特定的业务验证功能
3. 支持多个验证规则的组合
4. 验证器可以根据需要进行扩展和自定义 