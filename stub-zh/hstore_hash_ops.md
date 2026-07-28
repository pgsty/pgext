## 用法

来源：

- [官方上游 README](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/README.md)
- [官方扩展控制文件 (hstore_hash_ops.control)](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/hstore_hash_ops.control)
- [官方扩展 SQL (hstore_hash_ops--1.0.sql)](https://github.com/obartunov/hstore_ops/blob/8839ab7447e913f9109e94cffd38b78aa97d5505/hstore_hash_ops--1.0.sql)

`hstore_hash_ops` — Revived 非默认 GIN 操作类，已移植到当前 PostgreSQL 主分支（20devel）：当应用程序需要此特定数据库功能时，请使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION hstore_hash_ops;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `gin_compare_hstore_hash(int8, int8)` 是一个扩展函数，返回 `int4`。
- `gin_compare_hstore_pair(bytea, bytea)` 是一个扩展函数，返回 `int4`。
- `gin_compare_partial_hstore_hash(int8, int8, int2, internal)` 是一个扩展函数，返回 `int4`。
- `gin_consistent_hstore_hash(internal, int2, internal, int4, internal, internal)` 是一个扩展函数。
- `gin_consistent_hstore_pair(internal, int2, internal, int4, internal, internal)` 是一个扩展函数。
- `gin_extract_hstore_hash(internal, internal)` 是一个扩展函数，返回 `internal`。
- `gin_extract_hstore_pair(internal, internal)` 是一个扩展函数，返回 `internal`。
- `gin_extract_hstore_query_hash(internal, internal, int2, internal, internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `gin_extract_hstore_query_pair(internal, internal, int2, internal, internal, internal, internal)` 是一个扩展定义的操作类，返回 `internal`。
- `gin_hstore_hash_ops` 是一个扩展定义的操作类。
- `gin_hstore_pair_ops` 是一个扩展定义的操作类。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 请先安装确认的扩展依赖项：`hstore`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
