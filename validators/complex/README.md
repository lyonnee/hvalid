# Complex Validators

复杂验证器包，提供高级验证功能和验证器组合能力。这些验证器用于处理复杂的验证场景。

## 目录结构

- `logic/`: 逻辑组合验证器
  - `combination.go`: 验证器组合
  - `condition.go`: 条件验证
  - `logic.go`: 逻辑验证

- `async/`: 异步验证器
  - `async.go`: 异步验证
  - `retry.go`: 重试验证
  - `timeout.go`: 超时验证

- `chain/`: 链式验证器
  - `chain.go`: 验证器链
  - `transform.go`: 数据转换
  - `convert.go`: 类型转换

- 其他验证器
  - `aggregate.go`: 聚合验证
  - `batch.go`: 批量验证
  - `cache.go`: 缓存验证
  - `dependency.go`: 依赖验证

## 使用示例

```go
import "github.com/lyonnee/hvalid/validators/complex"

// 创建条件验证器
conditionValidator := complex.NewConditionValidator("field")

// 条件验证
err := conditionValidator.When(
    func(v int) bool { return v > 0 },
    func(v int) error { return nil },
)("value")
if err != nil {
    fmt.Printf("验证失败: %v\n", err)
}

// 异步验证
asyncValidator := complex.NewAsyncValidator("field")
err = asyncValidator.Validate(func(v string) error {
    // 异步验证逻辑
    return nil
})("value")
if err != nil {
    fmt.Printf("验证失败: %v\n", err)
}
```

## 注意事项

1. 复杂验证器提供了更高级的验证功能
2. 支持异步、重试、超时等特性
3. 可以组合多个验证器
4. 提供了数据转换和类型转换功能 