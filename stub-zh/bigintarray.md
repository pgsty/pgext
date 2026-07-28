## 用法

来源：

- [官方上游 README](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/README.md)
- [官方扩展控制文件 (bigintarray.control)](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/bigintarray.control)
- [官方扩展 SQL (bigintarray--1.0.sql)](https://github.com/pjungwir/bigintarray/blob/c0c6ba77e9b296009debe6c7acdbaded7c6a64a7/bigintarray--1.0.sql)

`bigintarray` — bigintarray 扩展提供了一维 bigints 数组（bigint[]）的相关函数、操作符和索引支持，行为与 PostgreSQL 内置的 intarray 扩展类似，但针对的是 bigints（8 字节整数）数组而非 integers（4 字节整数）数组。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION bigintarray;

-- Does the array contain both 1 and 2?
SELECT '{1,2,3}'::bigint[] @@ '1&2'::query_bigint;  -- true

-- Does the array contain 1 or 3?
SELECT '{1,2,3}'::bigint[] @@ '1|3'::query_bigint;  -- true

-- Does the array contain 1 and not 5?
SELECT '{1,2,3}'::bigint[] @@ '1&!5'::query_bigint;  -- true
```

在目标数据库中安装该扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bigintarray_del_elem(_int8, int8)` 是一个扩展函数并返回 `_int8`。
- `bigintarray_push_array(_int8, _int8)` 是一个扩展函数并返回 `_int8`。
- `bigintarray_push_elem(_int8, int8)` 是一个扩展函数并返回 `_int8`。
- `bigintset(int8)` 是一个扩展函数并返回 `_int8`。
- `bigintset_subtract(_int8, _int8)` 是一个扩展函数并返回 `_int8`。
- `bigintset_union_elem(_int8, int8)` 是一个扩展函数并返回 `_int8`。
- `boolop(_int8, query_bigint)` 是一个扩展函数。
- `bqarr_in(cstring)` 是一个扩展函数并返回 `query_bigint`。
- `bqarr_out(query_bigint)` 是一个扩展函数并返回 `cstring`。
- `g_bigint_compress(internal)` 是一个扩展函数并返回 `internal`。
- `g_bigint_consistent(internal,_int8,smallint,oid,internal)` 是一个扩展函数。
- `g_bigint_decompress(internal)` 是一个扩展函数并返回 `internal`。
- `g_bigint_options(internal)` 是一个扩展函数并返回 `void`。
- `g_bigint_penalty(internal,internal,internal)` 是一个扩展函数并返回 `internal`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将该扩展标记为可重定位。
- 控制文件将该扩展标记为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
