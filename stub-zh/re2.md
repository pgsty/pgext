## 用法

来源：[official README](https://github.com/ClickHouse/pg_re2/blob/main/README.md)、[official reference doc](https://github.com/ClickHouse/pg_re2/blob/main/doc/re2.md)、[v0.3.0 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.3.0)

`re2` 提供由 Google RE2 引擎驱动、与 ClickHouse 兼容的正则表达式函数。它同时暴露 `text` 和 `bytea` 重载，因此也可以搜索包含 `\\0` 字节的二进制数据。Pigsty 为 PostgreSQL 16-18 打包版本 `0.3.0`，而上游文档说明支持 PostgreSQL 13+。

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
- `re2extractallgroupsvertical(haystack, pattern) -> text[]|bytea[]`
- `re2extractallgroupshorizontal(haystack, pattern) -> text[]|bytea[]`
- `re2regexpquotemeta(haystack) -> text|bytea`
- `re2splitbyregexp(pattern, haystack, max_substrings default 0) -> text[]|bytea[]`
- `re2replaceregexpone(haystack, pattern, replacement) -> text|bytea`
- `re2replaceregexpall(haystack, pattern, replacement) -> text|bytea`
- `re2countmatches(...)` 和 `re2countmatchescaseinsensitive(...)`

```sql
SELECT re2extractallgroupsvertical('a=1 b=2', '(\\w)=(\\d)');
SELECT re2regexpquotemeta('a+b?');
SELECT re2splitbyregexp('\\s+', 'one two three', 2);
```

### 多模式匹配

`re2multimatch*` 系列既接受多个 pattern 参数，也接受 `VARIADIC` 数组：

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### 匹配语义

- 为了匹配 ClickHouse 行为，`.` 默认会匹配换行。
- 如果希望 `.` 不跨越换行，请在 pattern 前加上 `(?-s)`。
- 替换字符串支持 `\\0` 到 `\\9` 反向引用。

### 注意事项

- 上游要求构建/安装时系统中有 `re2` 库。
- Release `v0.3.0` 使用 SQL version `0.3`；从更早 minor release 替换扩展二进制后，执行 `ALTER EXTENSION re2 UPDATE TO '0.3'`。
- `re2splitbyregexp` 在 `v0.3.0` 中将参数顺序改为 `pattern, haystack[, max_substrings]`，与 ClickHouse 一致。早期 `0.2.0` build 使用的是 `haystack, pattern`。
- 上游将 patch releases 视为 binary-only，但 minor releases 可能需要 SQL upgrade scripts。
