## 用法

来源: [official README](https://github.com/ClickHouse/pg_re2/blob/main/README.md), [official reference doc](https://github.com/ClickHouse/pg_re2/blob/main/doc/re2.md), [v0.1.1 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.1.1)

`re2` 提供由 Google RE2 引擎驱动、与 ClickHouse 兼容的正则表达式函数。它同时暴露 `text` 和 `bytea` 重载，因此也可以搜索包含 `\\0` 字节的二进制数据。

```sql
CREATE EXTENSION re2;

SELECT re2match('hello world', 'h.*o');
SELECT re2extract('Order #123', '(\\d+)');
SELECT re2countmatches('a1 b2 c3', '\\d');
```

### 核心函数

- `re2match(haystack, pattern) -> boolean`
- `re2extract(haystack, pattern) -> text|bytea`
- `re2extractall(haystack, pattern) -> text[]|bytea[]`
- `re2regexpextract(haystack, pattern, index default 1) -> text|bytea`
- `re2extractgroups(haystack, pattern) -> text[]|bytea[]`
- `re2replaceregexpone(haystack, pattern, replacement) -> text|bytea`
- `re2replaceregexpall(haystack, pattern, replacement) -> text|bytea`
- `re2countmatches(...)` 和 `re2countmatchescaseinsensitive(...)`

### 多模式匹配

`re2multimatch*` 系列既接受多个 pattern 参数，也接受 `VARIADIC` 数组：

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### 匹配语义

- 为了匹配 ClickHouse 的行为，`.` 默认会匹配换行。
- 如果希望 `.` 不跨越换行，请在 pattern 前加上 `(?-s)`。
- 替换字符串支持 `\\0` 到 `\\9` 反向引用。

### 注意事项

- 构建或安装时需要系统 `re2` 库。
- `v0.1.1` 是仅涉及二进制构建的版本：它增加了 PostgreSQL 13+ 支持，并记录了多模式函数对 `VARIADIC` 的用法，但已有 `v0.1` SQL 安装不需要执行 `ALTER EXTENSION UPDATE`。
