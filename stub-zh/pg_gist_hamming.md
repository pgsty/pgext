## 用法

来源：

- [官方上游 README](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/README.md)
- [官方扩展控制文件 (pg_gist_hamming.control)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/old/pg_gist_hamming.control)
- [官方扩展 SQL (pg_gist_hamming--1.0.sql)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/old/pg_gist_hamming--1.0.sql)

`pg_gist_hamming` — 支持在 GiST 中索引常见数据类型。使用它来对应向量、模型或检索工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_gist_hamming;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `gbt_decompress(internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_compress(internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_consistent(internal,int8,int2,oid,internal)` 是一个扩展函数。
- `gbt_int8_distance(internal,int8,int2,oid,internal)` 是一个扩展函数并返回 `float8`。
- `gbt_int8_fetch(internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_hamming_distance(int8, int8)` 是一个扩展函数并返回 `int8`。
- `gbt_int8_penalty(internal,internal,internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_picksplit(internal, internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_same(gbtreekey16, gbtreekey16, internal)` 是一个扩展函数并返回 `internal`。
- `gbt_int8_union(internal, internal)` 是一个扩展函数并返回 `gbtreekey16`。
- `gbtreekey16_in(cstring)` 是一个扩展函数并返回 `gbtreekey16`。
- `gbtreekey16_out(gbtreekey16)` 是一个扩展函数并返回 `cstring`。
- `gbtreekey32_in(cstring)` 是一个扩展函数并返回 `gbtreekey32`。
- `gbtreekey32_out(gbtreekey32)` 是一个扩展函数并返回 `cstring`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
