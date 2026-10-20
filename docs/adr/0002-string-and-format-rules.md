# 字符串形态与格式规则不携带字段名

形态规则放在 package `hvalid`。格式规则放在 `validators/format`。两者都返回 `ValidatorFunc[T]`（`T ~string`），失败时只给出原因。字段名仍只在 `Validate` 上给一次。

`validators/primitive` 与 `validators/common` 里记住字段名的构造器是既有例外。新规则不模仿它们。

根包不引用 `validators`。格式包引用根包。不新增依赖，判定用标准库，不移植 govalidator 的正则。

空字符串对形态规则和格式规则都失败。空串本身由 `Blank` 判定。

## 形态规则

| 规则 | 通过条件 | 失败原因 |
| --- | --- | --- |
| `ASCII` | 非空，每个字节小于 128 | `must be ASCII` |
| `PrintableASCII` | 非空，每个字节在 0x20 到 0x7E | `must be printable ASCII` |
| `Alpha` | 非空，仅 ASCII 字母 | `must contain only ASCII letters` |
| `Alphanumeric` | 非空，仅 ASCII 字母和数字 | `must contain only ASCII letters and digits` |
| `Hex` | 非空，仅 `0-9a-fA-F` | `must be hexadecimal` |
| `Digits` | 非空，仅 ASCII `0-9` | `must contain only ASCII digits` |
| `Integer` | 可选 `+` 或 `-`，其后至少一位数字，整段匹配 | `must be an integer` |
| `Decimal` | 整段是十进制数，允许符号、小数、指数 | `must be a decimal number` |
| `UTFLetter` | 非空，每个符文都是字母 | `must contain only letters` |
| `UTFLetterNumeric` | 非空，每个符文是字母或数字 | `must contain only letters and numbers` |
| `UTFNumeric` | 非空，每个符文都是数字 | `must contain only numbers` |
| `UTFDigit` | 非空，每个符文都是十进制数字 | `must contain only Unicode digits` |
| `Lower` | 至少一个字母，每个字母都是小写 | `must be lower case` |
| `Upper` | 至少一个字母，每个字母都是大写 | `must be upper case` |
| `HasLower` | 至少一个小写字母 | `must contain a lower case letter` |
| `HasUpper` | 至少一个大写字母 | `must contain an upper case letter` |
| `HasSpace` | 至少一个空白符文 | `must contain whitespace` |
| `SpaceOnly` | 非空，每个符文都是空白 | `must contain only whitespace` |
| `Blank` | 空串 | `must be blank` |
| `NotBlank` | 非空串 | `must not be blank` |
| `OneOf` | 与允许列表中的某一项精确相等，区分大小写 | `must be one of: a, b` |
| `MinRunes` | 符文数不少于给定值 | `value rune length too short` |
| `MaxRunes` | 符文数不超过给定值 | `value rune length too long` |

`Integer` 接受 `0`、`-1`、`+1`，拒绝空串、`-`、`1.0`。`Decimal` 接受 `.5`、`1e-2`，拒绝 `1.`、`NaN`、`Inf`、`0x1`。`Lower` 与 `Upper` 允许数字和符号，因此 `a1` 通过，`1` 与空串失败。`MinRunes` 与 `MaxRunes` 按符文计数；`n < 0` 时不做特殊处理。

## 格式规则

| 规则 | 通过条件 | 失败原因 |
| --- | --- | --- |
| `UUID` | `8-4-4-4-12` 十六进制，版本半字节 1–8，RFC 变体 | `must be a UUID` |
| `ULID` | 26 位 Crockford Base32 | `must be a ULID` |
| `JSON` | 非空且是 JSON | `must be JSON` |
| `Base64` | 标准 Base64，填充正确，不含空白 | `must be base64` |
| `DataURI` | `data:` 前缀，逗号分隔；`;base64` 时载荷合法 | `must be a data URI` |
| `MAC` | `net.ParseMAC` 成功 | `must be a MAC address` |
| `DNSName` | RFC 1123 主机名 | `must be a DNS name` |
| `Host` | IP 或 DNS 名 | `must be a host` |
| `Port` | 十进制，值在 0 到 65535 | `must be a port` |
| `MongoID` | 恰好 24 位十六进制 | `must be a MongoDB object id` |
| `JWT` | 三段非空 Base64URL，前两段是 JSON 对象 | `must be a JWT` |
| `Semver` | 语义化版本 2.0.0，无 `v` 前缀 | `must be a semantic version` |
| `HexColor` | `#RGB`、`#RRGGBB` 或 `#RRGGBBAA` | `must be a hex color` |
| `RGB` | `rgb(r, g, b)`，分量 0–255 | `must be an RGB color` |
| `Latitude` | 十进制，-90 到 90 | `must be a latitude` |
| `Longitude` | 十进制，-180 到 180 | `must be a longitude` |
| `RFC3339` | RFC3339，可带小数秒 | `must be RFC3339` |
| `UnixTime` | 可选负号加数字，能放进 int64 | `must be a unix time` |
| `YYYYMMDD` | 8 位数字且是真实公历日期 | `must be a YYYYMMDD date` |
| `ISBN` | ISBN-10 或 ISBN-13，校验位正确 | `must be an ISBN` |
| `ISO3166Alpha2` | 现行 ISO 3166-1 alpha-2 | `must be an ISO 3166-1 alpha-2 code` |
| `ISO4217` | 现行 ISO 4217 字母代码 | `must be an ISO 4217 code` |

`UUID` 不接受 `urn:` 前缀和花括号。未传版本时接受 1–8；传入版本时半字节必须命中其中之一。`ULID` 拒绝 I、L、O、U，大小写不敏感。`Host` 不接受带方括号的 IPv6。`Port` 允许前导零，不允许符号。`JWT` 不验证签名。`Semver` 禁止数字部分的前导零。`ISBN` 去掉连字符和空格后再校验，ISBN-13 前缀为 978 或 979。国家码与货币码大小写不敏感，代码表以包内文件为准。
