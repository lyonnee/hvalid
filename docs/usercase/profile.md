# 资料字段走查

调用方有一份资料，逐项给出字段名并运行规则。规则本身不带字段名。某一项失败时，把 `Validate` 返回的错误原样留下，不再包一层。

| 字段 | 规则 | 通过样例 |
| --- | --- | --- |
| nickname | `Alpha`，`MinRunes(2)` | `Ann` |
| displayName | `NotBlank` | `Ann Lee` |
| country | `ISO3166Alpha2` | `CN` |
| currency | `ISO4217` | `CNY` |
| color | `HexColor` | `#336699` |
| userID | `UUID` | `550e8400-e29b-41d4-a716-446655440000` |
| isbn | `ISBN` | `978-0-306-40615-7` |
| updated | `RFC3339` | `2026-10-19T16:42:00Z` |
| port | `Port` | `8080` |
| host | `Host` | `example.com` |

失败走查把 nickname 换成 `A1`，displayName 换成空串。两项都失败，错误里分别出现 `nickname:` 与 `displayName:`，原因分别是 `must contain only ASCII letters` 与 `must not be blank`。
