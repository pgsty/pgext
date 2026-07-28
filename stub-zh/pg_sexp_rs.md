## 用法

来源：

- [官方上游 README](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/README.md)
- [官方扩展控制文件 (pg_sexp_rs.control)](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/rs/pg_sexp_rs.control)
- [官方实现源代码](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/rs/src/lib.rs)

`pg_sexp_rs` — PostgreSQL 扩展，用于添加对 s-表达式存储、查询和索引的支持。当应用程序数据需要这种类型或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_sexp_rs;

-- Atoms
SELECT 'hello'::sexp;           -- symbol
SELECT '42'::sexp;              -- integer
SELECT '3.14'::sexp;            -- float
SELECT '"hello world"'::sexp;   -- string
SELECT '()'::sexp;              -- nil

-- Lists
SELECT '(a b c)'::sexp;
SELECT '(define x 10)'::sexp;
SELECT '(lambda (x) (* x x))'::sexp;
```

在目标数据库中安装扩展，如果有可用示例，请运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `sexp_car` 是一个扩展函数。
- `sexp_cdr` 是一个扩展函数。
- `sexp_contains` 是一个扩展函数。
- `sexp_contains_key` 是一个扩展函数。
- `sexp_eq` 是一个扩展函数。
- `sexp_extract_keys` 是一个扩展函数。
- `sexp_extract_query_keys` 是一个扩展函数。
- `sexp_find` 是一个扩展函数。
- `sexp_gin_consistent_fn` 是一个扩展函数。
- `sexp_gin_extract_query_fn` 是一个扩展函数。
- `sexp_gin_extract_value_fn` 是一个扩展函数。
- `sexp_gin_triconsistent_fn` 是一个扩展函数。
- `sexp_hash` 是一个扩展函数。
- `sexp_hash_extended` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
