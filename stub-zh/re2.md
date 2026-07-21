## 用法

来源：

- [pg_re2 v0.4.1 README](https://github.com/ClickHouse/pg_re2/blob/v0.4.1/README.md)
- [SQL 参考](https://github.com/ClickHouse/pg_re2/blob/v0.4.1/doc/re2.md)
- [v0.4.0 发行说明](https://github.com/ClickHouse/pg_re2/releases/tag/v0.4.0)
- [v0.4.1 发行说明](https://github.com/ClickHouse/pg_re2/releases/tag/v0.4.1)

`re2` 提供与 ClickHouse 兼容、由 Google RE2 引擎实现的正则表达式函数。它同时提供 `text` 与 `bytea` 重载，因此也能搜索包含 `\\0` 字节的二进制数据。版本 `0.4.1` 还增加了索引辅助匹配，并可报告所链接的 RE2 版本。

```sql
CREATE EXTENSION re2;

SELECT re2match('hello world', 'h.*o');
SELECT re2extract('Order #123', '(\\d+)');
SELECT re2countmatches('a1 b2 c3', '\\d');
SELECT re2_version();
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
- `re2countmatches(...)` 与 `re2countmatchescaseinsensitive(...)`

```sql
SELECT re2extractallgroupsvertical('a=1 b=2', '(\\w)=(\\d)');
SELECT re2regexpquotemeta('a+b?');
SELECT re2splitbyregexp('\\s+', 'one two three', 2);
```

### 多模式匹配

`re2multimatch*` 函数族既可接受多个模式参数，也可接受 `VARIADIC` 数组：

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### 索引支持

版本 `0.4.0` 增加两种互补的索引路径：

```sql
-- Anchored constant patterns can use a normal btree prefix scan.
CREATE INDEX docs_body_btree ON docs (body);
SELECT * FROM docs WHERE re2match(body, '^order_2025');

-- The @~ operator can use the extension's GIN operator class.
CREATE INDEX docs_body_re2_gin ON docs USING gin (body gin_re2_ops);
SELECT * FROM docs WHERE body @~ 'timeout|denied';
```

扩展还为 RE2 谓词提供选择率估算。在 btree、GIN 与顺序扫描之间做选择前，请用代表性数据检查 `EXPLAIN`。

### 匹配语义

- 为匹配 ClickHouse 行为，`.` 默认可以匹配换行符。
- 如果不希望 `.` 跨越换行，请在模式前添加 `(?-s)`。
- 替换字符串支持 `\\0` 到 `\\9` 的反向引用。

### 注意事项

- 上游要求在构建/安装时存在系统 `re2` 库。
- `v0.4.x` 二进制使用 SQL 扩展版本 `0.4`；替换旧二进制后，如有待处理的升级，请运行 `ALTER EXTENSION re2 UPDATE TO '0.4'`。
- `v0.4.1` 修复了与缓存相关的释放后使用问题，并改善稳定模式与多模式匹配性能；应使用它而不是 `v0.4.0`。
- `re2splitbyregexp` 的参数顺序是 `pattern, haystack[, max_substrings]`；早于 `0.3.0` 的构建使用相反顺序。
- RE2 有意不支持模式中的反向引用和环视断言等功能；它的有界时间行为与 PostgreSQL 原生正则表达式引擎不同。
