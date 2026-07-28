## 用法

来源：

- [官方上游 README](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/README.md)
- [官方扩展控制文件 (falcon.control)](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/falcon/falcon.control)
- [官方扩展 SQL (falcon--1.0.sql)](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/falcon/falcon--1.0.sql)

`falcon` — FalconFS 是一个为 AI 工作负载设计的高性能分布式文件系统（DFS）。使用它来进行相应的分析或存储工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION falcon;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pg_catalog.falcon_acquire_hash_lock(IN path cstring, IN parentId bigint, IN lockmode bigint)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_build_shard_table(shard_count int)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_clear_all_data_func()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_clear_cached_relation_oid_func()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_clear_user_data_func()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_create_distributed_data_table()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_create_distributed_data_table_by_range_point(range_point int)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_create_kvmeta_table()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_create_slice_table()` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_delete_foreign_server(server_id int)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_drop_distributed_data_table_by_range_point(range_point int)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_foreign_server_test(mode cstring)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_insert_foreign_server(server_id int, server_name cstring, host cstring, port int, is_local bool, user_name cstring)` 是一个扩展函数，返回 `INTEGER`。
- `pg_catalog.falcon_meta_call_by_serialized_data(type int, count int, param bytea)` 是一个扩展函数，返回 `bytea`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
