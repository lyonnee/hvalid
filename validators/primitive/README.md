# Primitive Validators

字符串格式检查留在这个包里：`IsEmail`、`IsIPv4`、`IsIPv6`、`IsURL`。

长度、数值、切片、map、布尔、时间和相等这些规则在 package `hvalid`。

```go
import "github.com/lyonnee/hvalid/validators/primitive"

email := primitive.NewStringValidator("email")
err := email.IsEmail()("user@example.com")
```
