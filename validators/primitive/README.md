# Primitive Validators

基础验证器包，提供最基本的验证功能。这些验证器是构建其他验证器的基础组件。

## 包含的验证器

- `boolean.go`: 布尔值验证器
- `bytes.go`: 字节切片验证器
- `text.go`: 文本验证器
- `number.go`: 数字验证器
- `string.go`: 字符串验证器
- `time.go`: 时间验证器
- `map.go`: Map验证器
- `slice.go`: 切片验证器

## 使用示例

```go
import "github.com/lyonnee/hvalid/validators/primitive"

// 创建字符串验证器
stringValidator := primitive.NewStringValidator("name")

// 验证字符串长度
err := stringValidator.MinLen(3)("ab")
if err != nil {
    fmt.Printf("验证失败: %v\n", err)
}
```

## 注意事项

1. 所有验证器都实现了 `Validator` 接口
2. 验证器方法返回的是验证函数，需要传入值才能执行验证
3. 验证器支持链式调用 