# Primitive Validators

邮箱、IPv4、IPv6 和 URL 留在这个包里。构造器记住字段名，这是既有例外。

字符串形态规则在 package `hvalid`。UUID、ISBN、国家码等格式规则在 `validators/format`，不携带字段名。

长度、数值、切片、map、布尔、时间和相等这些规则也在 package `hvalid`。

```go
import "github.com/lyonnee/hvalid/validators/primitive"

email := primitive.NewStringValidator("email")
err := email.IsEmail()("user@example.com")
```
