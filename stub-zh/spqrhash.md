## 用法

来源：

- [官方上游 README](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/README.rst)
- [官方扩展控制文件 (spqrhash.control)](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/spqrhash.control)
- [官方扩展 SQL (spqrhash--1.1--1.2.sql)](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/sql/spqrhash--1.1--1.2.sql)

`spqrhash` — 该扩展提供了 SPQR 可以与 PG 一起使用的哈希函数。当 SQL 需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION spqrhash;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `spqrhash_city32(bytea)` 是一个扩展函数，返回 `int8`。
- `spqrhash_city32(id uuid)` 是一个扩展函数，返回 `int8`。
- `spqrhash_city32(int8)` 是一个扩展函数，返回 `int8`。
- `spqrhash_city32(int8[])` 是一个扩展函数，返回 `int8`。
- `spqrhash_city32(text)` 是一个扩展函数，返回 `int8`。
- `spqrhash_murmur3(bytea)` 是一个扩展函数，返回 `int8`。
- `spqrhash_murmur3(id uuid)` 是一个扩展函数，返回 `int8`。
- `spqrhash_murmur3(int8)` 是一个扩展函数，返回 `int8`。
- `spqrhash_murmur3(int8[])` 是一个扩展函数，返回 `int8`。
- `spqrhash_murmur3(text)` 是一个扩展函数，返回 `int8`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.2`。
- 控制文件标记该扩展为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
