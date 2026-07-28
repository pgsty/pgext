## 用法

来源：

- [官方上游 README](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/README.md)
- [官方扩展控制文件 (ecaz.control)](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/ecaz.control)
- [官方扩展 SQL (ecaz--0.1.0--0.1.1.sql)](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/ecaz--0.1.0--0.1.1.sql)

`ecaz` — Ecaz 是一个基于 Rust 的 PostgreSQL 扩展，用于高性能、高可扩展性的向量存储和检索。它支持广泛的量化和索引选项，而不是单一固定的架构。使用它来对应于向量、模型或检索工作流。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION ecaz;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `ec_distann_active_head_policy(index_regclass regclass)` 是一个扩展函数，返回 `TABLE`。
- `ec_distann_build_epoch_with_training(index_regclass regclass, epoch bigint, build_id uuid, training_relation regclass)` 是一个扩展函数，返回 `bytea`。
- `ec_spire_coordinator_index_shape_fingerprint(index_oid regclass)` 是一个扩展函数，返回 `text`。
- `ec_spire_coordinator_insert_shape_fingerprint(table_oid regclass)` 是一个扩展函数，返回 `text`。
- `ec_spire_register_placement_batch(index_oid oid, entries ec_spire_placement_entry[])` 是一个扩展函数，返回 `bigint`。
- `ec_spire_remote_catalog_drop_index_cleanup_event()` 是一个扩展函数，返回 `event_trigger`。
- `ec_spire_remote_index_shape_fingerprint(index_oid regclass)` 是一个扩展函数，返回 `text`。
- `ecvector(ecvector, integer, boolean)` 是一个扩展函数，返回 `ecvector`。
- `ecvector_from_bytea(bytea, integer, boolean)` 是一个扩展函数，返回 `ecvector`。
- `ecvector_from_real_array(real[], integer, boolean)` 是一个扩展函数，返回 `ecvector`。
- `ecvector_in(cstring, oid, integer)` 是一个扩展函数，返回 `ecvector`。
- `ecvector_inner_product(ecvector, ecvector)` 是一个扩展函数，返回 `float4`。
- `ecvector_negative_inner_product(ecvector, ecvector)` 是一个扩展函数，返回 `float4`。
- `ecvector_negative_query_inner_product(ecvector, real[])` 是一个扩展函数，返回 `float4`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `0.1.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 上游将该项目描述为一个概念验证。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
